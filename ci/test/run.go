package main

import (
	"fmt"
	"os"

	"github.com/vmware-tanzu/secrets-manager/ci/test/deploy"
	"github.com/vmware-tanzu/secrets-manager/ci/test/eval"
	"github.com/vmware-tanzu/secrets-manager/ci/test/state"
)

func sadCuddle(err error) {
	fmt.Println("Error running tests:", err.Error())
	os.Exit(1)
}

func try(fn func() error) {
	if err := fn(); err != nil {
		sadCuddle(err)
	}
}

func tryAll(fns ...func() error) {
	for _, fn := range fns {
		try(fn)
	}
}

func run() {
	tryAll(
		eval.SecretEncryption,
		state.Cleanup,
		deploy.WorkloadUsingSDK,
		eval.SecretRegistration,
		eval.SecretDeletion,
		eval.SecretRegistrationAppend,
		eval.SecretRegistrationJSONFormat,
		eval.SecretRegistrationYAMLFormat,
		state.Cleanup,
		deploy.WorkloadUsingSidecar,
		eval.SecretRegistrationSidecar,
		eval.SecretDeletionSidecar,
		eval.SecretRegistrationAppendSidecar,
		eval.SecretRegistrationJSONFormatSidecar,
		eval.SecretRegistrationYAMLFormatSidecar,
		state.Cleanup,
		eval.InitContainer,
		state.Cleanup,
	)
}
