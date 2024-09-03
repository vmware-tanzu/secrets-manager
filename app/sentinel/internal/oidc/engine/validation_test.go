/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_invalidInput(t *testing.T) {
	tests := []struct {
		name               string
		workloads          []string
		encrypt            bool
		serializedRootKeys string
		secret             string
		delete             bool
		want               bool
	}{
		{
			name:      "Valid input",
			workloads: []string{"workload1"},
			encrypt:   false,
			secret:    "mysecret",
			delete:    false,
			want:      true,
		},
		{
			name:      "Empty workloads",
			workloads: []string{},
			encrypt:   false,
			secret:    "mysecret",
			delete:    false,
			want:      false,
		},
		{
			name:      "Empty secret with delete false",
			workloads: []string{"workload1"},
			encrypt:   false,
			secret:    "",
			delete:    false,
			want:      false,
		},
		{
			name:      "Empty secret with delete true",
			workloads: []string{"workload1"},
			encrypt:   false,
			secret:    "",
			delete:    true,
			want:      true,
		},
		{
			name:               "Encrypt true with serializedRootKeys",
			workloads:          []string{"workload1"},
			encrypt:            true,
			serializedRootKeys: "validkeys",
			secret:             "mysecret",
			delete:             false,
			want:               true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isValidSecretModification(tt.workloads, tt.encrypt, tt.serializedRootKeys, tt.secret, tt.delete)
			assert.Equal(t, tt.want, got)
		})
	}
}
