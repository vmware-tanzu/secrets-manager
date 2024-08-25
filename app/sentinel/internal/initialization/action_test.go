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

//import (
//	"github.com/stretchr/testify/assert"
//	"testing"
//	"time"
//)
//
//type MockInitializer struct {
//}
//
//func TestDoSleep(t *testing.T) {
//	start := time.Now()
//	doSleep(100)
//	elapsed := time.Since(start)
//
//	assert.GreaterOrEqual(t, elapsed.Milliseconds(), int64(100))
//	assert.Less(t, elapsed.Milliseconds(), int64(150)) // Allow some margin for system delays
//}

//import (
//	"context"
//	"errors"
//	"fmt"
//	"github.com/agiledragon/gomonkey/v2"
//	"testing"
//	"time"
//
//	"github.com/stretchr/testify/assert"
//	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
//	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
//)
//
//func TestProcessCommandBlock(t *testing.T) {
//	ctx := context.Background()
//
//	defaultPatches := func() *gomonkey.Patches {
//		p := gomonkey.NewPatches()
//
//		// add more patches here if needed
//		// ...
//
//		return p
//	}
//
//	tests := []struct {
//		name      string
//		sc        entity.SentinelCommand
//		setupMock func() *gomonkey.Patches
//		wantErr   bool
//	}{
//		{
//			name: "Successful command processing",
//			sc: entity.SentinelCommand{
//				WorkloadIds: []string{"workload1"},
//				Secret:      "test-secret",
//			},
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				return p.ApplyFuncReturn(safe.Post, nil)
//			},
//			wantErr: false,
//		},
//		{
//			name: "Error in command processing",
//			sc: entity.SentinelCommand{
//				WorkloadIds: []string{"workload2"},
//				Secret:      "error-secret",
//			},
//			setupMock: func() *gomonkey.Patches {
//				p := defaultPatches()
//				return p.ApplyFuncReturn(safe.Post, errors.New("error"))
//			},
//			wantErr: true,
//		},
//	}
//
//	for _, tt := range tests {
//		fmt.Println(tt.name)
//		t.Run(tt.name, func(t *testing.T) {
//			m := tt.setupMock()
//			t.Cleanup(m.Reset)
//
//			err := processCommandBlock(ctx, tt.sc)
//			if tt.wantErr {
//				assert.Error(t, err)
//			} else {
//				assert.NoError(t, err)
//			}
//		})
//	}
//}
//
