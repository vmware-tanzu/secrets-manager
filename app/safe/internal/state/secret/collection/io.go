package collection

import (
	"errors"
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
