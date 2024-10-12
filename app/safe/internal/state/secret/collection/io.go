package collection

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"os"
	"strings"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/io"
	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/stats"
	f "github.com/vmware-tanzu/secrets-manager/core/constants/file"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func populateSecretsFromFileStore(cid string) error {
	root := env.DataPathForSafe()
	files, err := os.ReadDir(root)
	if err != nil {
		return errors.Join(
			err,
			errors.New("populateSecrets: problem reading secrets directory"),
		)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fn := file.Name()
		if strings.HasSuffix(fn, f.AgeBackupExtension) {
			continue
		}

		key := strings.Replace(fn, f.AgeExtension, "", 1)

		_, exists := Secrets.Load(key)
		if exists {
			continue
		}

		secretOnDisk, err := io.ReadFromDisk(key)
		if err != nil {
			log.ErrorLn(&cid,
				"populateSecrets: problem reading secret from disk:",
				err.Error())
			continue
		}
		if secretOnDisk != nil {
			stats.CurrentState.Increment(key, Secrets.Load)
			Secrets.Store(key, *secretOnDisk)
		}
	}

	return nil
}

func populateSecretsFromPostgresqlDb(cid string) error {
	if !io.PostgresReady() {
		return errors.New("populateSecretsFromPostgresqlDb: Database connection is not ready")
	}

	pg := io.DB()

	rows, err := pg.Query(`SELECT name, data FROM "vsecm-secrets"`)
	if err != nil {
		return errors.Join(
			err,
			errors.New("populateSecretsFromPostgresqlDb: problem querying secrets from database"),
		)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.ErrorLn(&cid, "populateSecretsFromPostgresqlDb: problem closing rows:",
				err.Error())
		}
	}(rows)

	for rows.Next() {
		var name string
		var encryptedData string

		err := rows.Scan(&name, &encryptedData)
		if err != nil {
			log.ErrorLn(&cid,
				"populateSecretsFromPostgresqlDb: problem scanning row:",
				err.Error())
			continue
		}

		_, exists := Secrets.Load(name)
		if exists {
			continue
		}

		// Decode the base64 encoded data
		encryptedBytes, err := base64.StdEncoding.DecodeString(encryptedData)
		if err != nil {
			log.ErrorLn(&cid,
				"populateSecretsFromPostgresqlDb: problem decoding base64 data:",
				err.Error())
			continue
		}

		// Decrypt the data
		var decryptedBytes []byte
		fipsMode := env.FipsCompliantModeForSafe()

		if fipsMode {
			decryptedBytes, err = crypto.DecryptBytesAes(encryptedBytes)
		} else {
			decryptedBytes, err = crypto.DecryptBytesAge(encryptedBytes)
		}

		if err != nil {
			log.ErrorLn(&cid,
				"populateSecretsFromPostgresqlDb: problem decrypting secret:",
				err.Error())
			continue
		}

		// Unmarshal the JSON data
		var secret entity.SecretStored
		err = json.Unmarshal(decryptedBytes, &secret)
		if err != nil {
			log.ErrorLn(&cid,
				"populateSecretsFromPostgresqlDb: problem unmarshaling secret:",
				err.Error())
			continue
		}

		stats.CurrentState.Increment(name, Secrets.Load)
		Secrets.Store(name, secret)
	}

	if err := rows.Err(); err != nil {
		return errors.Join(
			err,
			errors.New("populateSecretsFromPostgresqlDb: problem iterating over rows"),
		)
	}

	return nil
}
