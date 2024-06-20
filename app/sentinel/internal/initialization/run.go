/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"context"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"os"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/env"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// RunInitCommands reads and processes initialization commands from a file.
//
// This function is designed to execute an initial set of commands that are
// declaratively defined in a file. The end result will be as if an operator
// manually entered those commands using the Sentinel CLI.
//
// The function opens the file and reads it line by line using a `bufio.Scanner`.
// Each line is expected to be a command in a specific format, typically
// key-value pairs separated by a defined separator. Lines that do not conform
// to the expected format are ignored.
//
// Special handling is applied for commands that require sleeping (pause
// execution for a specified duration) or processing a block of commands.
//
// The function supports dynamic commands, which are defined in the
// 'entity.SentinelCommand' struct.
//
// Key commands include:
// - workload: (w) Sets the WorkloadIds field in the SentinelCommand.
// - namespace: (n) Sets the Namespaces field.
// - secret: (s) Sets the Secret field.
// - transformation: (t) Sets the Template field.
// - sleep: (sleep) Enables sleep mode and sets the SleepIntervalMs field.
//
// If the file cannot be opened, the function logs an informational message and
// returns early. Errors encountered while reading the file or closing it are
// logged as errors.
func RunInitCommands(ctx context.Context) {
	cid := ctx.Value(key.CorrelationId).(*string)

	waitInterval := env.InitCommandRunnerWaitBeforeExecIntervalForSentinel()
	time.Sleep(waitInterval)

	// Ensure that we can acquire a source before proceeding.
	source := ensureSourceAcquisition(ctx)

	// Now, we are sure that we can acquire a source.
	// Try to do a VSecM Safe API request with the source.
	ensureApiConnectivity(ctx, cid)

	// No need to proceed if initialization has been completed already.
	if initCommandsExecutedAlready(ctx, source) {
		log.TraceLn(cid, "RunInitCommands: executed already. exiting")
		return
	}

	log.TraceLn(cid, "RunInitCommands: starting the init flow")

	// Now we know that we can establish a connection to VSecM Safe
	// and execute API requests. So, we can safely run init commands.

	log.TraceLn(cid, "RunInitCommands: before getting the scanner")

	// Parse the commands file and execute the commands in it.
	file, scanner := commandFileScanner(cid)
	if file == nil {
		log.ErrorLn(cid, "file is nil, exiting")
		return
	}
	if scanner == nil {
		log.ErrorLn(cid, "scanner is nil, exiting")
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.ErrorLn(cid,
				"RunInitCommands: Error closing initialization file: ",
				err.Error(),
			)
		}
	}(file)

	log.TraceLn(cid, "RunInitCommands: before parsing commands file")

	parseCommandsFile(ctx, cid, scanner)

	log.TraceLn(cid, "RunInitCommands: before marking keystone")

	// Mark the keystone secret.
	success := markKeystone(ctx, cid)
	if !success {
		log.TraceLn(cid, "RunInitCommands: failed to mark keystone. exiting")

		// If we cannot set the keystone secret, better to retry everything.
		panic("RunInitCommands: failed to set keystone secret")
		return
	}

	// Wait before notifying Keystone. This way, if there are things that
	// take time to reconcile, they have a chance to do so.
	waitInterval = env.InitCommandRunnerWaitIntervalBeforeInitComplete()
	time.Sleep(waitInterval)

	// Everything is set up.
	log.InfoLn(cid, "RunInitCommands: keystone secret set successfully.")
}
