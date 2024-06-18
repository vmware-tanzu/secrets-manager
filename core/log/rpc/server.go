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
	"fmt"
	stdlib "log"
	"net"

	"google.golang.org/grpc"

	"github.com/vmware-tanzu/secrets-manager/core/log/rpc/generated"
)

// server struct implements the UnimplementedLogServiceServer interface generated
// by gRPC. It provides the functionality to handle log messages sent over gRPC
// by implementing the SendLog method.
type server struct {
	generated.UnimplementedLogServiceServer
}

// SendLog prints the log message contained in the request to the standard output.
// This method demonstrates a simple logging service over gRPC, where log messages
// from clients are received and printed on the server side.
//
// Parameters:
//   - ctx (context.Context): The context for the request, which allows for
//     cancellation and request scoping.
//   - in (*generated.LogRequest): The log request containing the message to
//     be logged.
//
// Returns:
//   - (*generated.LogResponse, error): Returns an empty LogResponse and nil
//     error as the operation is expected to succeed without any conditional
//     logic.
func (s *server) SendLog(ctx context.Context, in *generated.LogRequest,
) (*generated.LogResponse, error) {
	fmt.Printf("%s", in.Message)
	return &generated.LogResponse{}, nil
}

// CreateLogServer initializes and starts a gRPC server for logging services.
// It sets up a TCP listener on the address specified by SentinelLoggerUrl,
// registers the LogServiceServer, and starts serving incoming connections.
//
// Returns:
//   - A pointer to the initialized grpc.Server if the server starts successfully.
//   - Nil if there is an error in creating the listener or starting the server.
//
// Errors:
//   - If it fails to create the TCP listener, it logs the error and returns nil.
//   - If the server fails to start serving, it logs the error but still returns
//     the server instance, which might not be in a fully operational state.
//
// Usage:
//
//	server := logger.CreateLogServer()
//	if server == nil {
//	    fmt.Println("Failed to start log server")
//	} else {
//	    fmt.Println("Log server started successfully")
//	}
func CreateLogServer() *grpc.Server {
	lis, err := net.Listen("tcp", SentinelLoggerUrl())
	if err != nil {
		stdlib.Printf(
			"Logger.CreateLogServer error creating log server: %v\n", err)
		return nil
	}
	s := grpc.NewServer()

	generated.RegisterLogServiceServer(s, &server{})

	stdlib.Printf("Logger.CreateLogServer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		stdlib.Printf(
			"Logger.CreateLogServer failed to serve log server: %v\n", err)
		return s
	}

	return s
}
