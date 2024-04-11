package safe

import (
	"io"
	"net/http"

	log "github.com/vmware-tanzu/secrets-manager/core/log/rpc"
)

func respond(cid *string, r *http.Response) {
	if r == nil {
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.ErrorLn(cid, "Post: Problem closing request body.", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.ErrorLn(cid, "Post: Unable to read the response body from VSecM Safe.", err.Error())
		return
	}

	println("")
	println(string(body))
	println("")
}

func printEndpointError(cid *string, err error) {
	log.ErrorLn(cid, "Post: I am having problem generating VSecM Safe "+
		"secrets api endpoint URL.", err.Error())
}

func printPayloadError(cid *string, err error) {
	log.ErrorLn(cid, "Post: I am having problem generating the payload.", err.Error())
}
