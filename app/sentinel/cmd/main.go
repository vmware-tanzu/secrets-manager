/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware, Inc.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"os"
)

func parseList(parser *argparse.Parser) *bool {
	return parser.Flag("l", "list", &argparse.Options{
		Required: false, Default: false, Help: "lists all registered workloads.",
	})
}

func parseUseKubernetes(parser *argparse.Parser) *bool {
	return parser.Flag("k", "use-k8s", &argparse.Options{
		Required: false, Default: false,
		Help: "update an associated Kubernetes secret upon save. " +
			"Overrides VSECM_SAFE_USE_KUBERNETES_SECRETS.",
	})
}

func parseDeleteSecret(parser *argparse.Parser) *bool {
	return parser.Flag("d", "delete", &argparse.Options{
		Required: false, Default: false,
		Help: "delete the secret associated with the workload.",
	})
}

func parseAppendSecret(parser *argparse.Parser) *bool {
	return parser.Flag("a", "append", &argparse.Options{
		Required: false, Default: false,
		Help: "append the secret to the existing secret collection" +
			" associated with the workload.",
	})
}

func parseNamespace(parser *argparse.Parser) *string {
	return parser.String("n", "namespace", &argparse.Options{
		Required: false, Default: "default",
		Help: "the namespace of the Kubernetes Secret to create.",
	})
}

func parseInputKeys(parser *argparse.Parser) *string {
	return parser.String("i", "input-keys", &argparse.Options{
		Required: false,
		Help:     "A string containing the private and public Age keys and AES seed, each separated by '\\n'.",
	})
}

func parseBackingStore(parser *argparse.Parser) *string {
	return parser.String("b", "store", &argparse.Options{
		Required: false,
		Help: "backing store type (file|memory) (default: file). " +
			"Overrides VSECM_SAFE_BACKING_STORE.",
	})
}

func parseWorkload(parser *argparse.Parser) *string {
	return parser.String("w", "workload", &argparse.Options{
		Required: false,
		Help: "name of the workload (i.e. the '$name' segment of its " +
			"ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/…')).",
	})
}

func parseSecret(parser *argparse.Parser) *string {
	return parser.String("s", "secret", &argparse.Options{
		Required: false,
		Help:     "the secret to store for the workload.",
	})
}

func parseTemplate(parser *argparse.Parser) *string {
	return parser.String("t", "template", &argparse.Options{
		Required: false,
		Help:     "the template used to transform the secret stored.",
	})
}

func parseFormat(parser *argparse.Parser) *string {
	return parser.String("f", "format", &argparse.Options{
		Required: false,
		Help: "the format to display the secrets in." +
			" Has effect only when `-t` is provided. " +
			"Valid values: yaml, json, and none. Defaults to none.",
	})
}

func parseEncrypt(parser *argparse.Parser) *bool {
	return parser.Flag("e", "encrypt", &argparse.Options{
		Required: false, Default: false,
		Help: "returns an encrypted version of the secret if used with `-s`; " +
			"decrypts the secret before registering it to the workload if used " +
			"with `-s` and `-w`.",
	})
}

func printUsage(parser *argparse.Parser) {
	fmt.Print(parser.Usage("safe"))
}

func printWorkloadNameNeeded() {
	fmt.Println("Please provide a workload name.")
	fmt.Println("")
	fmt.Println("type `safe -h` (without backticks) and press return for help.")
	fmt.Println("")
}

func printSecretNeeded() {
	fmt.Println("Please provide a secret.")
	fmt.Println("")
	fmt.Println("type `safe -h` (without backticks) and press return for help.")
	fmt.Println("")
}

func doPost(workload *string, secret *string, namespace *string,
	backingStore *string, useKubernetes *bool, template *string, format *string,
	encrypt *bool, deleteSecret *bool, appendSecret *bool, inputKeys *string,
) {
	workloadP := ""
	if workload != nil {
		workloadP = *workload
	}

	secretP := ""
	if secret != nil {
		secretP = *secret
	}

	namespaceP := ""
	if namespace != nil {
		namespaceP = *namespace
	}

	backingStoreP := ""
	if backingStore != nil {
		backingStoreP = *backingStore
	}

	useK8sP := false
	if useKubernetes != nil {
		useK8sP = *useKubernetes
	}

	templateP := ""
	if template != nil {
		templateP = *template
	}

	formatP := ""
	if format != nil {
		formatP = *format
	}

	encryptP := false
	if encrypt != nil {
		encryptP = *encrypt
	}

	deleteP := false
	if deleteSecret != nil {
		deleteP = *deleteSecret
	}

	appendP := false
	if appendSecret != nil {
		appendP = *appendSecret
	}

	inputKeysP := ""
	if inputKeys != nil {
		inputKeysP = *inputKeys
	}

	safe.Post(
		workloadP, secretP, namespaceP, backingStoreP, useK8sP,
		templateP, formatP, encryptP, deleteP, appendP, inputKeysP,
	)
}

func main() {
	parser := argparse.NewParser("safe", "Assigns secrets to workloads.")

	list := parseList(parser)
	useKubernetes := parseUseKubernetes(parser)
	deleteSecret := parseDeleteSecret(parser)
	appendSecret := parseAppendSecret(parser)
	namespace := parseNamespace(parser)
	inputKeys := parseInputKeys(parser)
	backingStore := parseBackingStore(parser)
	workload := parseWorkload(parser)
	secret := parseSecret(parser)
	template := parseTemplate(parser)
	format := parseFormat(parser)
	encrypt := parseEncrypt(parser)

	err := parser.Parse(os.Args)
	if err != nil {
		printUsage(parser)
		return
	}

	if list != nil && *list == true {
		safe.Get()
		return
	}

	// You need to provide a workload name if you are not encrypting a secret,
	// or if you are not providing input keys.
	if (workload == nil || *workload == "") &&
		(encrypt == nil || !*encrypt) &&
		(inputKeys == nil || *inputKeys == "") {
		printWorkloadNameNeeded()
		return
	}

	// You need to provide a secret value if you are not deleting a secret,
	// or if you are not providing input keys.
	if (secret == nil || *secret == "") &&
		(deleteSecret == nil || !*deleteSecret) &&
		(inputKeys == nil || *inputKeys == "") {
		printSecretNeeded()
		return
	}

	if namespace == nil || *namespace == "" {
		*namespace = "default"
	}

	if inputKeys == nil || *inputKeys == "" {
		*inputKeys = ""
	}

	doPost(workload, secret, namespace, backingStore,
		useKubernetes, template, format, encrypt, deleteSecret, appendSecret,
		inputKeys,
	)
}
