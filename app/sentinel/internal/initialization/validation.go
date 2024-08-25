package initialization

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/constants/key"
	"github.com/vmware-tanzu/secrets-manager/lib/backoff"
)

func (i *Initializer) initCommandsExecutedAlready(
	ctx context.Context, src *workloadapi.X509Source,
) bool {
	cid := ctx.Value(key.CorrelationId).(*string)

	i.Logger.TraceLn(cid, "check:initCommandsExecutedAlready")

	initialized := false

	err := backoff.RetryExponential(
		"RunInitCommands:CheckConnectivity",
		func() error {
			var err error
			initialized, err = i.Safe.CheckInitialization(ctx, src)
			return err
		})

	if err == nil {
		return initialized
	}

	// I shouldn't be here.
	panic("RunInitCommands" +
		":initCommandsExecutedAlready: failed to check command initialization")
}
