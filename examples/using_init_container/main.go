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
	"os"
	"os/signal"
	"syscall"
	"time"
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
		fmt.Printf("My secret: '%s'.\n", os.Getenv("SECRET"))
		fmt.Printf("My creds: username:'%s' password:'%s'.\n",
			os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		)
		fmt.Println("")

		time.Sleep(5 * time.Second)
	}
}
