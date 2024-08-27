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
	"context"
	"os"
	"time"

	"github.com/spiffe/go-spiffe/v2/workloadapi"

	"github.com/vmware-tanzu/secrets-manager/app/sentinel/internal/safe"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/log/std"
	"github.com/vmware-tanzu/secrets-manager/core/spiffe"
)

type OSFileOpener struct{}

func (OSFileOpener) Open(name string) (*os.File, error) {
	return os.Open(name)
}

type EnvConfigReader struct{}

func (EnvConfigReader) InitCommandPathForSentinel() string {
	return env.InitCommandPathForSentinel()
}

func (EnvConfigReader) InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration {
	return env.InitCommandRunnerWaitBeforeExecIntervalForSentinel()
}

func (EnvConfigReader) InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration {
	return env.InitCommandRunnerWaitIntervalBeforeInitComplete()
}

func (EnvConfigReader) NamespaceForVSecMSystem() string {
	return env.NamespaceForVSecMSystem()
}

type StandardLogger struct{}

func (StandardLogger) InfoLn(correlationID *string, v ...interface{}) {
	std.InfoLn(correlationID, v...)
}
func (StandardLogger) ErrorLn(correlationID *string, v ...interface{}) {
	std.ErrorLn(correlationID, v...)
}
func (StandardLogger) TraceLn(correlationID *string, v ...interface{}) {
	std.TraceLn(correlationID, v...)
}
func (StandardLogger) WarnLn(correlationID *string, v ...interface{}) {
	std.WarnLn(correlationID, v...)
}
func (StandardLogger) FatalLn(correlationID *string, v ...interface{}) {
	std.FatalLn(correlationID, v...)
}

type SafeClient struct{}

func (SafeClient) Check(ctx context.Context, src *workloadapi.X509Source) error {
	return safe.Check(ctx, src)
}

func (SafeClient) CheckInitialization(ctx context.Context, src *workloadapi.X509Source) (bool, error) {
	return safe.CheckInitialization(ctx, src)
}

func (SafeClient) Post(ctx context.Context, sc data.SentinelCommand) error {
	return safe.Post(ctx, sc)
}

type SpiffeClient struct{}

func (SpiffeClient) AcquireSourceForSentinel(ctx context.Context) (*workloadapi.X509Source, bool) {
	return spiffe.AcquireSourceForSentinel(ctx)
}
