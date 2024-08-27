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

	"github.com/vmware-tanzu/secrets-manager/sdk/core/constants/env"
)

// PollIntervalForInitContainer returns the time interval between each poll in the
// Watch function. The interval is specified in milliseconds as the
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
