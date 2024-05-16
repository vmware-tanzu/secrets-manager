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

import "github.com/akamensky/argparse"

func parseList(parser *argparse.Parser) *bool {
	return parser.Flag("l", "list", &argparse.Options{
		Required: false, Help: "lists all registered workloads",
	})
}

func parseDeleteSecret(parser *argparse.Parser) *bool {
	return parser.Flag("d", "delete", &argparse.Options{
		Required: false, Default: false,
		Help: "delete the secret associated with the workload",
	})
}

func parseAppendSecret(parser *argparse.Parser) *bool {
	return parser.Flag("a", "append", &argparse.Options{
		Required: false, Default: false,
		Help: "append the secret to the existing secret collection" +
			" associated with the workload",
	})
}

func parseNamespaces(parser *argparse.Parser) *[]string {
	return parser.StringList("n", "namespace", &argparse.Options{
		Required: false, Default: []string{"default"},
		Help: "the namespaces of the workloads or Kubernetes secrets",
	})
}

func parseInputKeys(parser *argparse.Parser) *string {
	return parser.String("i", "input-keys", &argparse.Options{
		Required: false,
		Help: "A string containing the private and public " +
			"Age keys and AES seed, each separated by '\\n'",
	})
}

func parseWorkload(parser *argparse.Parser) *[]string {
	return parser.StringList("w", "workload", &argparse.Options{
		Required: false,
		Help: "name of the workload (i.e. the '$name' segment of its " +
			"ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/...'))",
	})
}

func parseSecret(parser *argparse.Parser) *string {
	return parser.String("s", "secret", &argparse.Options{
		Required: false,
		Help:     "the secret to store for the workload",
	})
}

func parseTemplate(parser *argparse.Parser) *string {
	return parser.String("t", "template", &argparse.Options{
		Required: false,
		Help:     "the template used to transform the secret stored",
	})
}

func parseFormat(parser *argparse.Parser) *string {
	return parser.String("f", "format", &argparse.Options{
		Required: false,
		Help: "the format to display the secrets in." +
			" Has effect only when `-t` is provided. " +
			"Valid values: yaml, json, and none. Defaults to none",
	})
}

func parseEncrypt(parser *argparse.Parser) *bool {
	return parser.Flag("e", "encrypt", &argparse.Options{
		Required: false, Default: false,
		Help: "returns an encrypted version of the secret if used with " +
			"`-s`; decrypts the secret before registering it to the " +
			"workload if used with `-s` and `-w`",
	})
}

func parseExpires(parser *argparse.Parser) *string {
	return parser.String("E", "exp", &argparse.Options{
		Required: false, Default: "never",
		Help: "is the expiration date of the secret",
	})
}

func parseNotBefore(parser *argparse.Parser) *string {
	return parser.String("N", "nbf", &argparse.Options{
		Required: false, Default: "now",
		Help: "secret is not valid before this time",
	})
}
