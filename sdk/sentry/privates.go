/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentry

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"os"
)

func saveData(data string) error {
	path := env.SidecarSecretsPath()

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error saving data")
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		return errors.Wrap(err, "error saving data")
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing writer")
	}

	return nil
}

func fetchSecrets() error {
	r, eFetch := Fetch()

	// VMware Secrets Manager Safe was successfully queried, but no secrets found.
	// This means someone has deleted the secret. We cannot let
	// the workload linger with the existing secret, so we remove
	// it from the workload too.
	//
	// If the user wants a more fine-tuned control for this case,
	// that is: if the user wants to keep the existing secret even
	// if it has been deleted from VMware Secrets Manager Safe, then the user should
	// use VSecM SDK directly, instead of using VMware Secrets Manager Sidecar.
	if eFetch == ErrSecretNotFound {
		return saveData("")
	}

	v := r.Data
	if v == "" {
		return nil
	}
	return saveData(v)
}
