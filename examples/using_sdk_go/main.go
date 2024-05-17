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
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vmware-tanzu/secrets-manager/sdk/sentry" // <- SDK
)

func main() {
	go func() {
		// Block the process from exiting, but also be graceful and honor the
		// termination signals that may come from the orchestrator.
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		select {
		case e := <-s:
			fmt.Println(e)
			panic("bye cruel world!")
		}
	}()

	for {
		log.Println("fetch")
		d, err := sentry.Fetch()

		if err != nil {
			fmt.Println("Failed to read the secrets file. Will retry in 5 seconds...")
			fmt.Println(err.Error())
			time.Sleep(5 * time.Second)
			continue
		}

		if d.Data == "" {
			fmt.Println("no secret yet... will check again later.")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Printf(
			"secret: updated: %s, created: %s, value: %s\n",
			d.Updated, d.Created, d.Data,
		)
		time.Sleep(5 * time.Second)
	}
}
