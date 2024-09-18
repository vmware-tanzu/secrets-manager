/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package env

import (
	"strconv"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/constants/env"
)

// PollIntervalForInitContainer returns the time interval between each poll in
// the Watch function. The interval is specified in milliseconds as the
// VSECM_INIT_CONTAINER_POLL_INTERVAL environment variable.  If the environment
// variable is not set or is not a valid integer value, the function returns the
// default interval of 5000 milliseconds.
func PollIntervalForInitContainer() time.Duration {
	p := env.Value(env.VSecMInitContainerPollInterval)
	d, _ := strconv.Atoi(string(env.VSecMInitContainerPollIntervalDefault))
	if p == "" {
		p = string(env.VSecMInitContainerPollIntervalDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}

// WaitBeforeExitForInitContainer retrieves the wait time before exit for an
// init container. The duration is determined by the environment variable
// "VSECM_INIT_CONTAINER_WAIT_BEFORE_EXIT" and defaults to zero if the variable
// is not set or cannot be parsed.
//
// The environment variable is expected to be an integer value representing the
// wait time in milliseconds. If parsing fails, the function will return 0
// milliseconds.
//
// Returns:
//
//	time.Duration: The wait time before exit, in milliseconds.
func WaitBeforeExitForInitContainer() time.Duration {
	p := env.Value(env.VSecMInitContainerWaitBeforeExit)
	d, _ := strconv.Atoi(string(env.VSecMInitContainerWaitBeforeExitDefault))
	if p == "" {
		p = string(env.VSecMInitContainerWaitBeforeExitDefault)
	}

	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		i = int64(d)
		return time.Duration(i) * time.Millisecond
	}

	return time.Duration(i) * time.Millisecond
}
