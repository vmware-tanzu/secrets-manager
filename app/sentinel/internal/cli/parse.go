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

// ParseList adds a flag to the parser for listing all registered workloads.
// It returns a pointer to a boolean that will be set to true if the flag is
// used.
func ParseList(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.List),
		string(sentinel.ListExp), &argparse.Options{
			Required: false, Help: "lists all registered workloads",
		})
}

// ParseDeleteSecret adds a flag to the parser for deleting a secret associated
// with a workload.
// It returns a pointer to a boolean that will be set to true if the flag is
// used.
func ParseDeleteSecret(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.Remove),
		string(sentinel.RemoveExp), &argparse.Options{
			Required: false, Default: false,
			Help: "delete the secret associated with the workload",
		})
}

// ParseAppendSecret adds a flag to the parser for appending a secret to an
// existing secret collection.
// It returns a pointer to a boolean that will be set to true if the flag is
// used.
func ParseAppendSecret(parser *argparse.Parser) *bool {
	return parser.Flag(
		string(sentinel.Join),
		string(sentinel.JoinExp), &argparse.Options{
			Required: false, Default: false,
			Help: "append the secret to the existing secret collection" +
				" associated with the workload",
		})
}

// ParseNamespaces adds a string list argument to the parser for specifying
// namespaces.
// It returns a pointer to a slice of strings containing the specified
// namespaces.
func ParseNamespaces(parser *argparse.Parser) *[]string {
	return parser.StringList(
		string(sentinel.Namespace),
		string(sentinel.NamespaceExp), &argparse.Options{
			Required: false, Default: []string{"default"},
			Help: "the namespaces of the workloads or Kubernetes secrets",
		})
}

// ParseInputKeys adds a string argument to the parser for inputting keys.
// It returns a pointer to a string containing the input keys.
func ParseInputKeys(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Keys),
		string(sentinel.KeysExp), &argparse.Options{
			Required: false,
			Help: "A string containing the private and public " +
				"Age keys and AES seed, each separated by '\\n'",
		})
}

// ParseWorkload adds a string list argument to the parser for specifying
// workload names.
// It returns a pointer to a slice of strings containing the workload names.
func ParseWorkload(parser *argparse.Parser) *[]string {
	return parser.StringList(
		string(sentinel.Workload),
		string(sentinel.WorkloadExp), &argparse.Options{
			Required: false,
			Help:     "names of the workloads",
		})
}

// ParseSecret adds a string argument to the parser for specifying a secret.
// It returns a pointer to a string containing the secret.
func ParseSecret(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Secret),
		string(sentinel.SecretExp), &argparse.Options{
			Required: false,
			Help:     "the secret to store for the workload",
		})
}

// ParseTemplate adds a string argument to the parser for specifying a
// transformation template.
// It returns a pointer to a string containing the template.
func ParseTemplate(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Transformation),
		string(sentinel.TransformationExp), &argparse.Options{
			Required: false,
			Help:     "the template used to transform the secret stored",
		})
}

// ParseFormat adds a string argument to the parser for specifying the output
// format.
// It returns a pointer to a string containing the format.
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

// ParseEncrypt adds a flag to the parser for encrypting or decrypting secrets.
// It returns a pointer to a boolean that will be set to true if the flag is
// used.
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

// ParseExpires adds a string argument to the parser for specifying the
// expiration date of a secret.
// It returns a pointer to a string containing the expiration date.
func ParseExpires(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.Expires),
		string(sentinel.ExpiresExp), &argparse.Options{
			Required: false, Default: "never",
			Help: "is the expiration date of the secret",
		})
}

// ParseNotBefore adds a string argument to the parser for specifying the
// start date of a secret's validity.
// It returns a pointer to a string containing the start date.
func ParseNotBefore(parser *argparse.Parser) *string {
	return parser.String(
		string(sentinel.NotBefore),
		string(sentinel.NotBeforeExp), &argparse.Options{
			Required: false, Default: "now",
			Help: "secret is not valid before this time",
		})
}
