/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package probe

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ok(t *testing.T) {
	type args struct {
		w   *httptest.ResponseRecorder
		in1 *http.Request
	}
	tests := []struct {
		name       string
		setup      func()
		args       args
		statusWant int
		bodyWant   string
	}{
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
			},
			statusWant: http.StatusOK,
			bodyWant:   "OK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			ok(tt.args.w, tt.args.in1)
			if status := tt.args.w.Code; status != tt.statusWant {
				t.Errorf("ok status = %v, want %v", status, tt.statusWant)
			} else if body := tt.args.w.Body.String(); body != tt.bodyWant {
				t.Errorf("ok body = %v, want %v", body, tt.bodyWant)
			}
		})
	}
}
