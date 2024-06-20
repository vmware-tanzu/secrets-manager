/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package cli

import (
	"github.com/akamensky/argparse"

	"github.com/vmware-tanzu/secrets-manager/core/constants/sentinel"
)

func ParseList(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.List),
		string(sentinel.ListExp), &argparse.Options{
			Required: false, Help: "lists all registered workloads",
		})
}

func ParseDeleteSecret(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.Remove),
		string(sentinel.RemoveExp), &argparse.Options{
			Required: false, Default: false,
			Help: "delete the secret associated with the workload",
		})
}

func ParseAppendSecret(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.Join),
		string(sentinel.JoinExp), &argparse.Options{
			Required: false, Default: false,
			Help: "append the secret to the existing secret collection" +
				" associated with the workload",
		})
}

func ParseNamespaces(parser *argparse.Parser) *[]string {
	return parser.StringList(
		string(sentinel.Namespace),
		string(sentinel.NamespaceExp), &argparse.Options{
			Required: false, Default: []string{"default"},
			Help: "the namespaces of the workloads or Kubernetes secrets",
		})
}

func ParseInputKeys(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Keys),
		string(sentinel.KeysExp), &argparse.Options{
			Required: false,
			Help: "A string containing the private and public " +
				"Age keys and AES seed, each separated by '\\n'",
		})
}

func ParseWorkload(parser *argparse.Parser) *[]string {
	return parser.StringList(
		string(sentinel.Workload),
		string(sentinel.WorkloadExp), &argparse.Options{
			Required: false,
			Help: "name of the workload (i.e. the '$name' segment of its " +
				"ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/...'))",
		})
}

func ParseSecret(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Secret),
		string(sentinel.SecretExp), &argparse.Options{
			Required: false,
			Help:     "the secret to store for the workload",
		})
}

func ParseTemplate(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Transformation),
		string(sentinel.TransformationExp), &argparse.Options{
			Required: false,
			Help:     "the template used to transform the secret stored",
		})
}

func ParseFormat(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Format),
		string(sentinel.FormatExp), &argparse.Options{
			Required: false,
			Help: "the format to display the secrets in." +
				" Has effect only when `-t` is provided. " +
				"Valid values: yaml, json, and none. Defaults to none",
		})
}

func ParseEncrypt(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.Encrypt),
		string(sentinel.EncryptExp), &argparse.Options{
			Required: false, Default: false,
			Help: "returns an encrypted version of the secret if used with " +
				"`-s`; decrypts the secret before registering it to the " +
				"workload if used with `-s` and `-w`",
		})
}

func ParseExpires(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Expires),
		string(sentinel.ExpiresExp), &argparse.Options{
			Required: false, Default: "never",
			Help: "is the expiration date of the secret",
		})
}

func ParseNotBefore(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.NotBefore),
		string(sentinel.NotBeforeExp), &argparse.Options{
			Required: false, Default: "now",
			Help: "secret is not valid before this time",
		})
}
