package main

import (
	"context"
	"fmt"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/vmware-tanzu/secrets-manager/core/env"
	"github.com/vmware-tanzu/secrets-manager/core/validation"
	s "github.com/vmware-tanzu/secrets-manager/lib/spiffe"
	"io"
	"net/http"

	"github.com/vmware-tanzu/secrets-manager/core/crypto"
	log "github.com/vmware-tanzu/secrets-manager/core/log/std"
)

func fallback(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	log.DebugLn(&cid, "Handler: route mismatch:", r.RequestURI)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.WarnLn(&cid, "Problem writing response:", err.Error())
	}
}

func success(
	cid string, r *http.Request, w http.ResponseWriter,
) {
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "OK "+r.URL.Path)
	if err != nil {
		log.WarnLn(&cid, "Problem writing response:", err.Error())
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cid := crypto.Id()

	source, err := workloadapi.NewX509Source(
		ctx,
		workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)

	if err != nil {
		log.FatalLn(&cid, "Unable to fetch X.509 Bundle", err.Error())
		return
	}

	if source == nil {
		log.FatalLn(&cid, "Could not find source")
		return
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(&cid, "Unable to get X.509 SVID from source bundle", err.Error())
		return
	}

	if svid == nil {
		log.FatalLn(&cid, "Could not find SVID in source bundle")
		return
	}

	svidId := svid.ID
	if !validation.IsRelayServer(svidId.String()) {
		log.FatalLn(
			&cid,
			"SpiffeId check: RelayServer:bootstrap: I don't know you, and it's crazy:",
			svidId.String(),
		)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cid := crypto.Id()

		validation.EnsureRelayServer(source)

		id, err := s.IdFromRequest(r)

		if err != nil {
			log.WarnLn(&cid, "Handler: blocking insecure svid", id, err)

			fallback(cid, r, w)

			return
		}

		sid := s.IdAsString(r)

		p := r.URL.Path
		m := r.Method
		log.DebugLn(
			&cid,
			"Handler: got svid:", sid, "path", p, "method", m)

		success(cid, r, w)
	})

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsRelayClient(id.String()) {
			return nil
		}

		return fmt.Errorf(
			"TLS Config: I don't know you, and it's crazy '%s'", id.String(),
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      env.TlsPort(),
		TLSConfig: tlsConfig,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.FatalLn(&cid, "Failed to listen and serve", err.Error())
	}
}
