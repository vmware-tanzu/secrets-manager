/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package safe

import (
	"fmt"
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
		log.ErrorLn(cid,
			"Post: Unable to read the response body from VSecM Safe.",
			err.Error())
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
