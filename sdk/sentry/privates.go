/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentry

import (
	"bufio"
	"errors"
	"os"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

func saveData(data string) error {
	path := env.SecretsPathForSidecar()

	f, err := os.Create(path)
	if err != nil {
		return errors.Join(
			err,
			errors.New("error saving data"),
		)
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		return errors.Join(
			err,
			errors.New("error saving data"),
		)
	}

	err = w.Flush()
	if err != nil {
		return errors.Join(
			err,
			errors.New("error flushing writer"),
		)
	}

	return nil
}

func fetchSecrets() error {
	r, eFetch := Fetch()

	// VSecM Safe was successfully queried, but no secrets found.
	// This means someone has deleted the secret. We cannot let
	// the workload linger with the existing secret, so we remove
	// it from the workload too.
	//
	// If the user wants a more fine-tuned control for this case,
	// that is: if the user wants to keep the existing secret even
	// if it has been deleted from VSecM Safe, then the user should
	// use VSecM SDK directly, instead of using VSecM Sidecar.
	if errors.Is(eFetch, ErrSecretNotFound) {
		return saveData("")
	}

	v := r.Data
	if v == "" {
		return nil
	}
	return saveData(v)
}
