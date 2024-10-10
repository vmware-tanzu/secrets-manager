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

// TODO: obviously there is a need for cleanup; once things start to work
// as expected, move the codes to where they should belong.

// TODO: move me to a proper place.

// TODO: by design, VSecM Safe will not use more than one backing store
// (create an ADR for that).
// This means, there is a chicken-and-the-egg problem for persisting the
// internal VSecM Safe configuration.
//
// For postgres backing store, VSecM Safe should keep its initial config
// in memory until the database is there; and then it should save it to
// the database, too.

// TODO: we should check for the existence of the table in postgres and
// log an error if it's not there.

// TODO: when postgres mode vsecm safe shall be read-only (except for config update)
// until it is initialized. once initialized, it should save its config to postgres too
// and then it should be readwrite.

// TODO: we need documentation for this postgres store feature. (and also a demo recording)

// TODO: it's best block requests when the db is not ready yet (in postgres mode)
// because otherwise, the initCommand will retry in exponential backoff and
// eventually give up.
// or the keystone secret will not be persisted although keystone will
// be informed that safe is ready.

type SafeConfig struct {
	Config struct {
		BackingStore   string `json:"backingStore"`
		DataSourceName string `json:"dataSourceName"`
	} `json:"config"`
}
