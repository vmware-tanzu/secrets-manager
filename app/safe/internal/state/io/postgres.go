/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package io

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sync"
	"sync/atomic"

	_ "github.com/lib/pq"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

var (
	db     atomic.Pointer[sql.DB]
	initMu sync.Mutex
)

// InitDB initializes the database connection
func InitDB(dataSourceName string) error {
	initMu.Lock()
	defer initMu.Unlock()

	// Check if db is already initialized
	if db.Load() != nil {
		return nil
	}

	newDB, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}

	if err := newDB.Ping(); err != nil {
		_ = newDB.Close()
		return err
	}

	db.Store(newDB)
	return nil
}

func PostgresReady() bool {
	currentDB := db.Load()
	if currentDB == nil {
		return false
	}
	return currentDB.Ping() == nil
}

// DB returns the current database connection
func DB() *sql.DB {
	return db.Load()
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

	err = backoff.RetryExponential("PersistToPostgres", func() error {
		pg := DB()

		_, err := pg.Exec(
			`INSERT INTO "vsecm-secrets" (name, data) 
VALUES ($1, $2) ON CONFLICT (name) DO UPDATE SET data = $2`,
			secret.Name, encryptedData)
		return err
	})

	if err != nil {
		errChan <- errors.Join(err, errors.New("PersistToPostgres: Failed to persist secret to database"))
		log.ErrorLn(&cid, "PersistToPostgres: Error persisting secret to database:", err.Error())
		return
	}

	log.TraceLn(&cid, "PersistToPostgres: Secret persisted to database successfully")
}
