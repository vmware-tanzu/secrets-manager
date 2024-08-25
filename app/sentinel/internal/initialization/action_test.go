/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package initialization

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestInitializer_doSleep(t *testing.T) {
	tests := []struct {
		name    string
		seconds int
	}{
		{
			name:    "Sleep for 100 milliseconds",
			seconds: 1,
		},
		{
			name:    "Sleep for 200 milliseconds",
			seconds: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tn := time.Now()
			i := Initializer{}
			i.doSleep(tt.seconds)

			assert.GreaterOrEqual(t, time.Since(tn).Milliseconds(), int64(tt.seconds))
		})
	}
}
