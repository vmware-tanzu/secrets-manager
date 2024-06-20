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
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/akamensky/argparse"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/cli"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/core/constants/sentinel"
	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

func main() {
	id := crypto.Id()

	parser := argparse.NewParser(
		sentinel.CmdName,
		"Assigns secrets to workloads.",
	)

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(),
			key.CorrelationId, &id),
	)
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case <-c:
			fmt.Println("Operation was cancelled.")
			// It is okay to cancel a cancelled context.
			cancel()
		}
	}()

	list := cli.ParseList(parser)
	deleteSecret := cli.ParseDeleteSecret(parser)
	appendSecret := cli.ParseAppendSecret(parser)
	namespaces := cli.ParseNamespaces(parser)
	inputKeys := cli.ParseInputKeys(parser)
	workloadIds := cli.ParseWorkload(parser)
	secret := cli.ParseSecret(parser)
	template := cli.ParseTemplate(parser)
	format := cli.ParseFormat(parser)
	encrypt := cli.ParseEncrypt(parser)
	notBefore := cli.ParseNotBefore(parser)
	expires := cli.ParseExpires(parser)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println()
		cli.PrintUsage(parser)
		return
	}

	if *list {
		if *encrypt {
			err = safe.Get(ctx, true)
			if err != nil {
				fmt.Println("Error getting from VSecM Safe:", err.Error())
				return
			}

			return
		}

		err = safe.Get(ctx, false)
		if err != nil {
			fmt.Println("Error getting from VSecM Safe:", err.Error())
			return
		}

		return
	}

	if *namespaces == nil || len(*namespaces) == 0 {
		*namespaces = []string{string(env.Default)}
	}

	if cli.InputValidationFailure(
		workloadIds, encrypt, inputKeys, secret, deleteSecret,
	) {
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
		fmt.Println("Error posting to VSecM Safe:", err.Error())
		return
	}
}
