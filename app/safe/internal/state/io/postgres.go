package io

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"

	_ "github.com/lib/pq" // PostgreSQL driver
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

// PersistToPostgres saves a given secret to the Postgres database
func PersistToPostgres(secret entity.SecretStored, errChan chan<- error) {
	cid := secret.Meta.CorrelationId

	log.TraceLn(&cid, "PersistToPostgres: Persisting secret to database")

	// Serialize the secret to JSON
	jsonData, err := json.Marshal(secret)
	if err != nil {
		errChan <- errors.Join(err, errors.New("PersistToPostgres: Failed to marshal secret"))
		log.ErrorLn(&cid, "PersistToPostgres: Error marshaling secret:", err.Error())
		return
	}

	// Encrypt the JSON data
	var encryptedData string
	fipsMode := env.FipsCompliantModeForSafe()

	if fipsMode {
		encryptedBytes, err := crypto.EncryptBytesAes(jsonData)
		if err != nil {
			errChan <- errors.Join(err, errors.New("PersistToPostgres: Failed to encrypt secret with AES"))
			log.ErrorLn(&cid, "PersistToPostgres: Error encrypting secret with AES:", err.Error())
			return
		}
		encryptedData = base64.StdEncoding.EncodeToString(encryptedBytes)
	} else {
		encryptedBytes, err := crypto.EncryptBytesAge(jsonData)
		if err != nil {
			errChan <- errors.Join(err, errors.New("PersistToPostgres: Failed to encrypt secret with Age"))
			log.ErrorLn(&cid, "PersistToPostgres: Error encrypting secret with Age:", err.Error())
			return
		}
		encryptedData = base64.StdEncoding.EncodeToString(encryptedBytes)
	}

	// TODO: maybe retry with an exponential backoff
	// Persist the encrypted data to the database
	_, err = db.Exec(
		"INSERT INTO secrets (name, data) VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET data = $2",
		secret.Name, encryptedData)

	if err != nil {
		errChan <- errors.Join(err, errors.New("PersistToPostgres: Failed to persist secret to database"))
		log.ErrorLn(&cid, "PersistToPostgres: Error persisting secret to database:", err.Error())
		return
	}

	log.TraceLn(&cid, "PersistToPostgres: Secret persisted to database successfully")
}
