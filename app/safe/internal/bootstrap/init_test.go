/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package bootstrap

import "testing"

func Test_completeInitialization(t *testing.T) {
	id := "test_id"
	type args struct {
		correlationId *string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-1",
			args: args{
				correlationId: &id,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			completeInitialization(tt.args.correlationId)
		})
	}
}
