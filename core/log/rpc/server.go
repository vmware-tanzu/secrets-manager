/*
|    Protect your secrets, protect your sensitive data.
:    Explore VMware Secrets Manager docs at https://vsecm.com/
</
<>/  keep your secrets… secret
>/
<>/' Copyright 2023–present VMware Secrets Manager contributors.
>/'  SPDX-License-Identifier: BSD-2-Clause
*/

package rpc

import (
	"context"
	"fmt"
	"github.com/vmware-tanzu/secrets-manager/core/log/rpc/generated"
	"google.golang.org/grpc"
	stdlib "log"
	"net"
)

type server struct {
	generated.UnimplementedLogServiceServer
}

func (s *server) SendLog(ctx context.Context, in *generated.LogRequest) (*generated.LogResponse, error) {
	fmt.Printf("%s", in.Message)
	return &generated.LogResponse{}, nil
}

func CreateLogServer() *grpc.Server {
	lis, err := net.Listen("tcp", SentinelLoggerUrl())
	if err != nil {
		stdlib.Printf("Logger.CreateLogServer error creating log server: %v\n", err)
		return nil
	}
	s := grpc.NewServer()

	generated.RegisterLogServiceServer(s, &server{})

	stdlib.Printf("Logger.CreateLogServer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		stdlib.Printf("Logger.CreateLogServer failed to serve log server: %v\n", err)
		return s
	}

	return s
}
