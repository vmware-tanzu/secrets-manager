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
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vmware-tanzu/secrets-manager/core/env"
)

// CreateLiveness sets up and starts an HTTP server on the port specified by
// env.ProbeLivenessPort() to serve as a liveness probe for the application.
// The server listens for requests at the root path ("/") and responds with an
// "ok" message. If there is an error starting the server, the function logs
// a fatal message and returns.
func CreateLiveness() chan bool {
	ready := make(chan bool)

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", ok)
		err := http.ListenAndServe(env.ProbeLivenessPort(), mux)
		if err != nil {
			log.Fatalf("error creating liveness probe: %s", err.Error())
			return
		}
	}()

	go func() {
		for {
			resp, err := http.Get(fmt.Sprintf(
				"http://localhost%s/", env.ProbeLivenessPort()))
			if err == nil && resp.StatusCode == http.StatusOK {
				ready <- true
				return
			}
			time.Sleep(100 * time.Millisecond) // Wait before retrying
		}
	}()

	return ready
}

// CreateReadiness sets up and starts an HTTP server on the port specified by
// env.ProbeReadinessPort() to serve as a readiness probe for the application.
// The server listens for requests at the root path ("/") and responds with an
// "ok" message. If there is an error starting the server, the function logs
// a fatal message and returns.
func CreateReadiness() chan bool {
	ready := make(chan bool)

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", ok)
		err := http.ListenAndServe(env.ProbeReadinessPort(), mux)
		if err != nil {
			log.Fatalf("error creating readiness probe: %s", err.Error())
			return
		}
	}()

	go func() {
		for {
			resp, err := http.Get(fmt.Sprintf(
				"http://localhost%s/", env.ProbeReadinessPort()))
			if err == nil && resp.StatusCode == http.StatusOK {
				ready <- true
				return
			}
			time.Sleep(100 * time.Millisecond) // Wait before retrying
		}
	}()

	return ready
}
