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
	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go func() {
		// Block the process from exiting, but also be graceful and honor the
		// termination signals that may come from the orchestrator.
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		select {
		case e := <-s:
			println(e)
			panic("bye cruel world!")
		}
	}()

	// Fetch the secret from the VSecM Safe.
	d, err := sentry.Fetch()
	if err != nil {
		println("Failed to fetch the secrets. Try again later.")
		println(err.Error())
		return
	}

	if d.Data == "" {
		println("No secret yet... Try again later.")
		return
	}

	// d.Data is a collection of Secrets.
	println(d.Data)
}
