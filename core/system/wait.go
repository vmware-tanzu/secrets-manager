/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package system

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// KeepAlive keeps the system alive until a `SIGINT` or `SIGTERM` comes to
// the progress. It does that by opening up a channel and keeping it open
// until a termination signal comes.
//
// Make sure you run it on the main thread (NOT in a goroutine) for it to
// take effect.
func KeepAlive() {
	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	select {
	case e := <-s:
		fmt.Println(e)
		panic("bye cruel world!")
	}
}
