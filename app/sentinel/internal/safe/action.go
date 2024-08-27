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
	"bytes"
	"errors"
	"net/http"
)

func doDelete(cid *string, client *http.Client, p string, md []byte) error {
	req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
	if err != nil {
		return errors.Join(
			err,
			errors.New("post:Delete: Problem connecting"+
				" to VSecM Safe API endpoint URL"),
		)
	}

	req.Header.Set("Content-Type", "application/json")

	r, err := client.Do(req)
	if err != nil {
		return errors.Join(
			err,
			errors.New("post:Delete: Problem connecting"+
				" to VSecM Safe API endpoint URL"),
		)
	}

	respond(cid, r)
	return nil
}

func doPost(cid *string, client *http.Client, p string, md []byte) error {
	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		return errors.Join(
			err,
			errors.New("post: Problem connecting to VSecM Safe API endpoint URL"),
		)
	}
	respond(cid, r)
	return nil
}
