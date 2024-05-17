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
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vmware-tanzu/secrets-manager/sdk/sentry"
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

	// Fetch the secret from the VSecM Safe.
	d, err := sentry.Fetch()
	if err != nil {
		fmt.Println("Failed to fetch the secrets. Try again later.")
		fmt.Println(err.Error())
		return
	}

	if d.Data == "" {
		fmt.Println("No secret yet... Try again later.")
		return
	}

	// Check if d.Data is a JSON array
	if string(d.Data[0]) == "[" {
		// Convert the array into a slice of strings
		var dataSlice []string
		err = json.Unmarshal([]byte(d.Data), &dataSlice)
		if err != nil {
			fmt.Println("Failed to unmarshal the data into a slice of strings. Check the data format.")
			fmt.Println(err.Error())
			return
		}

		// Concatenate all members of the slice into one large string
		concatString := ""
		for _, s := range dataSlice {
			concatString += s
		}

		// Base64 decode the string
		decodedString, err := base64.StdEncoding.DecodeString(concatString)
		if err != nil {
			fmt.Println("Failed to decode the base64 string.")
			fmt.Println(err.Error())
			fmt.Println("Raw data:")
			fmt.Println(d.Data)
			return
		}

		// Print the result
		fmt.Println(string(decodedString))
	} else {
		// d.Data is a collection of Secrets.
		fmt.Println(d.Data)
	}
}
