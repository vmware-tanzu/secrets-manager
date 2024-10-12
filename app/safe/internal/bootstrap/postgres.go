package bootstrap

import (
	"context"
	"encoding/json"
	"time"

	"github.com/vmware-tanzu/secrets-manager/app/safe/internal/state/secret/collection"
	"github.com/vmware-tanzu/secrets-manager/core/entity/v1/data"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

// PollForConfig continuously polls for the VSecM Safe internal configuration.
//
// This function attempts to read the configuration from a secret. It will keep
// polling until a valid configuration is found or the context is cancelled.
//
// Parameters:
//   - ctx: A context.Context for cancellation and timeout control.
//   - id: A correlation ID for logging purposes.
//
// Returns:
//   - *data.VSecMSafeInternalConfig: A pointer to the parsed VSecM Safe internal
//     configuration.
//   - error: An error if the context is cancelled or if there's an issue parsing
//     the configuration.
//
// The function will log informational messages about its progress and any
// errors encountered.
// It sleeps for 5 seconds between each polling attempt.
func PollForConfig(id string, ctx context.Context,
) (*data.VSecMSafeInternalConfig, error) {
	log.InfoLn(&id, "Will poll for VSecM Safe internal configuration...")

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			vSecMSafeInternalConfig, err := collection.ReadSecret(id, "vsecm-safe")

			if err != nil {
				log.InfoLn(&id, "Failed to load VSecM Safe internal configuration", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}

			if vSecMSafeInternalConfig == nil {
				log.InfoLn(&id, "VSecM Safe internal configuration not found")
				time.Sleep(5 * time.Second)
				continue
			}

			if len(vSecMSafeInternalConfig.Values) == 0 {
				log.InfoLn(&id, "VSecM Safe internal configuration is empty")
				time.Sleep(5 * time.Second)
				continue
			}

			var safeConfig data.VSecMSafeInternalConfig

			err = json.Unmarshal(
				[]byte(vSecMSafeInternalConfig.Values[0]), &safeConfig,
			)

			if err != nil {
				log.InfoLn(&id, "Failed to parse VSecM Safe internal configuration", err.Error())
				time.Sleep(5 * time.Second)
				continue
			}

			return &safeConfig, nil
		}
	}
}
