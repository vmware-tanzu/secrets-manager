/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"context"
	"fmt"
	entity "github.com/vmware-tanzu/secrets-manager/core/entity/data/v1"
	"os"
	"os/signal"
	"syscall"

	"github.com/akamensky/argparse"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
)

func main() {
	parser := argparse.NewParser("safe", "Assigns secrets to workloads.")

	list := parseList(parser)
	useKubernetes := parseUseKubernetes(parser)
	deleteSecret := parseDeleteSecret(parser)
	appendSecret := parseAppendSecret(parser)
	namespace := parseNamespace(parser)
	inputKeys := parseInputKeys(parser)
	backingStore := parseBackingStore(parser)
	workloadId := parseWorkload(parser)
	secret := parseSecret(parser)
	template := parseTemplate(parser)
	format := parseFormat(parser)
	encrypt := parseEncrypt(parser)
	notBefore := parseNotBefore(parser)
	expires := parseExpires(parser)

	err := parser.Parse(os.Args)
	if err != nil {
		printUsage(parser)
		return
	}

	if *list {
		if *encrypt {
			safe.Get(true)
			return
		}
		safe.Get(false)
		return
	}

	if *namespace == "" {
		*namespace = "default"
	}

	if inputValidationFailure(workloadId, encrypt, inputKeys, secret, deleteSecret) {
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case <-c:
			fmt.Println("Operation was cancelled.")
			cancel()
		}
	}()

	safe.Post(ctx, entity.SentinelCommand{
		WorkloadId:    *workloadId,
		Secret:        *secret,
		Namespace:     *namespace,
		BackingStore:  *backingStore,
		UseKubernetes: *useKubernetes,
		Template:      *template,
		Format:        *format,
		Encrypt:       *encrypt,
		DeleteSecret:  *deleteSecret,
		AppendSecret:  *appendSecret,
		InputKeys:     *inputKeys,
		NotBefore:     *notBefore,
		Expires:       *expires,
	})
}
