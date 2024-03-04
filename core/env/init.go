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
	"os"
	"strconv"
	"time"
)

// PollIntervalForInitContainer returns the time interval between each poll in the
// Watch function. The interval is specified in milliseconds as the
// VSECM_INIT_CONTAINER_POLL_INTERVAL environment variable.  If the environment
// variable is not set or is not a valid integer value, the function returns the
// default interval of 5000 milliseconds.
func PollIntervalForInitContainer() time.Duration {
	p := os.Getenv("VSECM_INIT_CONTAINER_POLL_INTERVAL")
	if p == "" {
		p = "5000"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 5000 * time.Millisecond
	}
	return time.Duration(i) * time.Millisecond
}
