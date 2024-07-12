/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package sentinel

// SecretRequest encapsulates a VSecM Safe REST command payload.
type SecretRequest struct {
	Workloads          []string `json:"workload"`
	Secret             string   `json:"secret"`
	Namespaces         []string `json:"namespaces,omitempty"`
	Encrypt            bool     `json:"encrypt,omitempty"`
	Delete             bool     `json:"delete,omitempty"`
	Append             bool     `json:"append,omitempty"`
	List               bool     `json:"list,omitempty"`
	Template           string   `json:"template,omitempty"`
	Format             string   `json:"format,omitempty"`
	SerializedRootKeys string   `json:"root-keys,omitempty"`
	NotBefore          string   `json:"nbf,omitempty"`
	Expires            string   `json:"exp,omitempty"`
}
