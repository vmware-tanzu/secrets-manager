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
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var id string = "test_id"

func TestOSFileOpener_Open(t *testing.T) {
	f, err := os.Open("./impl.go")
	assert.NoError(t, err)
	type args struct {
		name string
	}
	tests := []struct {
		name      string
		o         OSFileOpener
		args      args
		want      *os.File
		wantErr   bool
		afterTest func(*os.File, *os.File)
	}{
		{
			name:    "successful file open test",
			o:       OSFileOpener{},
			args:    args{name: "./impl.go"},
			want:    f,
			wantErr: false,
			afterTest: func(actualFile, expectedFile *os.File) {
				if actualFile != nil {
					defer actualFile.Close()
				}
				if expectedFile != nil {
					defer expectedFile.Close()
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OSFileOpener{}
			got, err := o.Open(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("OSFileOpener.Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name(), tt.want.Name()) {
				t.Errorf("OSFileOpener.Open() = %v, want %v", got, tt.want)
			}
			if tt.afterTest != nil {
				tt.afterTest(got, f)
			}
		})
	}
}

func TestEnvConfigReader_InitCommandPathForSentinel(t *testing.T) {
	tests := []struct {
		name      string
		testSetup func()
		want      string
		afterTest func()
	}{
		{
			name: "custom path from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_PATH", "/samplePath")
				assert.NoError(t, err)
			},
			want: "/samplePath",
			afterTest: func() {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_PATH")
				assert.NoError(t, err)
			},
		},
		{
			name: "default path",
			want: "/opt/vsecm-sentinel/init/data",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testSetup != nil {
				tt.testSetup()
			}
			if tt.afterTest != nil {
				defer tt.afterTest()
			}
			e := EnvConfigReader{}
			if got := e.InitCommandPathForSentinel(); got != tt.want {
				t.Errorf("EnvConfigReader.InitCommandPathForSentinel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvConfigReader_InitCommandRunnerWaitBeforeExecIntervalForSentinel(t *testing.T) {
	tests := []struct {
		name      string
		testSetup func()
		want      time.Duration
		afterTest func()
	}{
		{
			name: "custom value from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC", "5")
				assert.NoError(t, err)
			},
			want: time.Millisecond * 5,
			afterTest: func() {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC")
				assert.NoError(t, err)
			},
		},
		{
			name: "invalid custom value from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC", "invalid_time")
				assert.NoError(t, err)
			},
			want: 0,
			afterTest: func() {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_BEFORE_EXEC")
				assert.NoError(t, err)
			},
		},
		{
			name: "default value",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testSetup != nil {
				tt.testSetup()
			}
			if tt.afterTest != nil {
				defer tt.afterTest()
			}
			e := EnvConfigReader{}
			if got := e.InitCommandRunnerWaitBeforeExecIntervalForSentinel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvConfigReader.InitCommandRunnerWaitBeforeExecIntervalForSentinel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvConfigReader_InitCommandRunnerWaitIntervalBeforeInitComplete(t *testing.T) {
	tests := []struct {
		name      string
		testSetup func()
		want      time.Duration
		afterTest func()
	}{
		{
			name: "custom value from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE", "100")
				assert.NoError(t, err)
			},
			want: time.Millisecond * 100,
			afterTest: func() {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE")
				assert.NoError(t, err)
			},
		},
		{
			name: "invalid custom value from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE", "invalid_time")
				assert.NoError(t, err)
			},
			want: 0,
			afterTest: func() {
				err := os.Unsetenv("VSECM_SENTINEL_INIT_COMMAND_WAIT_AFTER_INIT_COMPLETE")
				assert.NoError(t, err)
			},
		},
		{
			name: "default value",
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testSetup != nil {
				tt.testSetup()
			}
			if tt.afterTest != nil {
				defer tt.afterTest()
			}
			e := EnvConfigReader{}
			if got := e.InitCommandRunnerWaitIntervalBeforeInitComplete(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EnvConfigReader.InitCommandRunnerWaitIntervalBeforeInitComplete() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvConfigReader_NamespaceForVSecMSystem(t *testing.T) {
	tests := []struct {
		name      string
		testSetup func()
		want      string
		afterTest func()
	}{
		{
			name: "custom value from env variable",
			testSetup: func() {
				err := os.Setenv("VSECM_NAMESPACE_SYSTEM", "vsecm-custom")
				assert.NoError(t, err)
			},
			want: "vsecm-custom",
			afterTest: func() {
				err := os.Unsetenv("VSECM_NAMESPACE_SYSTEM")
				assert.NoError(t, err)
			},
		},
		{
			name: "default value",
			want: "vsecm-system",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.testSetup != nil {
				tt.testSetup()
			}
			if tt.afterTest != nil {
				defer tt.afterTest()
			}
			e := EnvConfigReader{}
			if got := e.NamespaceForVSecMSystem(); got != tt.want {
				t.Errorf("EnvConfigReader.NamespaceForVSecMSystem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStandardLogger_InfoLn(t *testing.T) {
	type args struct {
		correlationID *string
		v             []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "info log test",
			args: args{
				correlationID: &id,
				v:             []interface{}{"test string 1", "test string 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StandardLogger{}
			s.InfoLn(tt.args.correlationID, tt.args.v...)
		})
	}
}

func TestStandardLogger_ErrorLn(t *testing.T) {
	type args struct {
		correlationID *string
		v             []interface{}
	}
	tests := []struct {
		name string
		s    StandardLogger
		args args
	}{
		{
			name: "log error test",
			args: args{
				correlationID: &id,
				v:             []interface{}{"test string 1", "test string 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StandardLogger{}
			s.ErrorLn(tt.args.correlationID, tt.args.v...)
		})
	}
}

func TestStandardLogger_TraceLn(t *testing.T) {
	type args struct {
		correlationID *string
		v             []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "log trace test",
			args: args{
				correlationID: &id,
				v:             []interface{}{"test string 1", "test string 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StandardLogger{}
			s.TraceLn(tt.args.correlationID, tt.args.v...)
		})
	}
}

func TestStandardLogger_WarnLn(t *testing.T) {
	type args struct {
		correlationID *string
		v             []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "log warn test",
			args: args{
				correlationID: &id,
				v:             []interface{}{"test string 1", "test string 2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := StandardLogger{}
			s.WarnLn(tt.args.correlationID, tt.args.v...)
		})
	}
}
