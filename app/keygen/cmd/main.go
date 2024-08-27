/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"os"

	"github.com/vmware-tanzu/secrets-manager/app/keygen/internal"
	e "github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func main() {
	id := crypto.Id()

	// Print the diagnostic information about the environment.
	log.PrintEnvironmentInfo(&id, []string{
		string(e.AppVersion),
		string(e.VSecMLogLevel),
		string(e.VSecMKeygenDecrypt),
	})

	if env.KeyGenDecrypt() {
		// This is a Kubernetes Secret, mounted as a file.
		keyPath := env.RootKeyPathForKeyGen()

		if _, err := os.Stat(keyPath); os.IsNotExist(err) {
			log.FatalLn(&id,
				"CreateRootKey: Secret key not mounted at", keyPath)
			return
		}

		data, err := os.ReadFile(keyPath)
		if err != nil {
			log.FatalLn(&id,
				"CreateRootKey: Error reading file:", err.Error())
			return
		}

		// Root key needs to be committed to memory for VSecM Keygen to be able
		// to decrypt the secrets.
		secret := string(data)
		crypto.SetRootKeyInMemory(secret)

		internal.PrintDecryptedKeys()
		return
	}

	internal.PrintGeneratedKeys()
}
