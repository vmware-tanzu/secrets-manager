package rest

import (
	"github.com/vmware-tanzu/secrets-manager/app/sentinel/rest/core"
	"log"
	"net/http"
)

func RunRestServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/secrets", core.HandleSecrets)
	log.Println("VSecM Server started at :8085")
	log.Fatal(http.ListenAndServe(":8085", mux))
}
