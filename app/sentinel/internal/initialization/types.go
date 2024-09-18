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

	entity "github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
)

// FileOpener is an interface for the file operations that the Sentinel
// needs to perform.
type FileOpener interface {
	Open(name string) (*os.File, error)
}

// EnvReader is an interface for the environment variables that the Sentinel
// needs to read.
type EnvReader interface {
	InitCommandPathForSentinel() string
	InitCommandRunnerWaitBeforeExecIntervalForSentinel() time.Duration
	InitCommandRunnerWaitIntervalBeforeInitComplete() time.Duration
	NamespaceForVSecMSystem() string
}

// Logger is an interface for the logging operations that the Sentinel
// needs to perform.
type Logger interface {
	InfoLn(correlationID *string, v ...any)
	ErrorLn(correlationID *string, v ...any)
	TraceLn(correlationID *string, v ...any)
	WarnLn(correlationID *string, v ...any)
	FatalLn(correlationID *string, v ...any)
}

// SafeOps is an interface for the Safe operations that the Sentinel needs
// to perform.
type SafeOps interface {
	Check(ctx context.Context, src *workloadapi.X509Source) error
	CheckInitialization(ctx context.Context,
		src *workloadapi.X509Source) (bool, error)
	Post(ctx context.Context, sc entity.SentinelCommand) error
}

// SpiffeOps is an interface for the Spiffe operations that the Sentinel needs
// to perform.
type SpiffeOps interface {
	AcquireSourceForSentinel(
		ctx context.Context) (*workloadapi.X509Source, bool)
}

type Initializer struct {
	FileOpener FileOpener
	EnvReader  EnvReader
	Logger     Logger
	Safe       SafeOps
	Spiffe     SpiffeOps
}

func NewInitializer(
	fileOpener FileOpener,
	envReader EnvReader,
	logger Logger,
	safe SafeOps,
	spiffe SpiffeOps,
) *Initializer {
	return &Initializer{
		FileOpener: fileOpener,
		EnvReader:  envReader,
		Logger:     logger,
		Safe:       safe,
		Spiffe:     spiffe,
	}
}
