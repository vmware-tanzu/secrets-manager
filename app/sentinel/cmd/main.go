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
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/akamensky/argparse"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

const defaultNs = "default"

func main() {
	id := crypto.Id()

	parser := argparse.NewParser("safe", "Assigns secrets to workloads.")

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
	)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case <-c:
			println("Operation was cancelled.")
			// It is okay to cancel a cancelled context.
			cancel()
		}
	}()

	list := parseList(parser)
	deleteSecret := parseDeleteSecret(parser)
	appendSecret := parseAppendSecret(parser)
	namespaces := parseNamespaces(parser)
	inputKeys := parseInputKeys(parser)
	workloadIds := parseWorkload(parser)
	secret := parseSecret(parser)
	template := parseTemplate(parser)
	format := parseFormat(parser)
	encrypt := parseEncrypt(parser)
	notBefore := parseNotBefore(parser)
	expires := parseExpires(parser)

	err := parser.Parse(os.Args)
	if err != nil {
		println(err.Error())
		println()
		printUsage(parser)
		return
	}

	if *list {
		if *encrypt {
			err = safe.Get(ctx, true)
			if err != nil {
				println("Error getting from VSecM Safe:", err.Error())
				return
			}

			return
		}

		err = safe.Get(ctx, false)
		if err != nil {
			println("Error getting from VSecM Safe:", err.Error())
			return
		}

		return
	}

	if *namespaces == nil || len(*namespaces) == 0 {
		*namespaces = []string{defaultNs}
	}

	if inputValidationFailure(workloadIds, encrypt, inputKeys, secret, deleteSecret) {
		return
	}

	err = safe.Post(ctx, entity.SentinelCommand{
		WorkloadIds:        *workloadIds,
		Secret:             *secret,
		Namespaces:         *namespaces,
		Template:           *template,
		Format:             *format,
		Encrypt:            *encrypt,
		DeleteSecret:       *deleteSecret,
		AppendSecret:       *appendSecret,
		SerializedRootKeys: *inputKeys,
		NotBefore:          *notBefore,
		Expires:            *expires,
	})

	if err != nil {
		println("Error posting to VSecM Safe:", err.Error())
		return
	}
}
