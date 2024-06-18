/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets... secret
>/
<>/' Copyright 2023-present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	"context"
	stdlib "log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/vmware-tanzu/secrets-manager/core/log/rpc/generated"
)

func log(message string) {
	conn, err := grpc.Dial(
		SentinelLoggerUrl(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		stdlib.Printf("Logger.log could not connect to server: %v\n", err)
		return
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			stdlib.Printf("Logger.log could not close connection: %v\n", err)
		}
	}(conn)

	c := generated.NewLogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SendLog(ctx, &generated.LogRequest{Message: message})
	if err != nil {
		stdlib.Printf("Logger.log could not send message: %v\n", err)
		return
	}
}
