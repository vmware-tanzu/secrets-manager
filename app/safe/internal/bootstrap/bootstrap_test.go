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

import (
	"testing"
)

var id string = "test_id"

func TestChannelsToMonitor_Size(t *testing.T) {
	type fields struct {
		AcquiredSvid  <-chan bool
		UpdatedSecret <-chan bool
		ServerStarted <-chan bool
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test-1",
			fields: fields{
				AcquiredSvid:  make(<-chan bool),
				UpdatedSecret: make(<-chan bool),
				ServerStarted: make(<-chan bool),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChannelsToMonitor{
				AcquiredSvid:  tt.fields.AcquiredSvid,
				UpdatedSecret: tt.fields.UpdatedSecret,
				ServerStarted: tt.fields.ServerStarted,
			}
			if got := c.Size(); got != tt.want {
				t.Errorf("ChannelsToMonitor.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
