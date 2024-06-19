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
	"net"
	"os"
	"strings"
	"testing"
	"time"

	"google.golang.org/grpc"

	"github.com/vmware-tanzu/secrets-manager/core/log/rpc/generated"
)

var cid = "test-correlation-id"

type MockLogServiceServer struct {
	generated.UnimplementedLogServiceServer
	ReceivedMessage string
}

func (s *MockLogServiceServer) SendLog(ctx context.Context,
	in *generated.LogRequest) (*generated.LogResponse, error) {
	s.ReceivedMessage = in.Message
	return &generated.LogResponse{}, nil
}

func TestCreateLogServer(t *testing.T) {
	// Start the server in a goroutine
	go func() {
		server := CreateLogServer()
		if server == nil {
			t.Error("Failed to create gRPC server")
		}
	}()

	// Give the server a moment to start
	time.Sleep(time.Second)

	// Attempt to connect to the server
	_, err := net.Dial("tcp", SentinelLoggerUrl())
	if err != nil {
		t.Errorf("Failed to connect to the server: %v", err)
	}
}

func TestSendLogMessage(t *testing.T) {
	server := &MockLogServiceServer{}
	grpcServer := grpc.NewServer()
	generated.RegisterLogServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("error creating log server: %v", err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(lis)

	_ = os.Setenv("VSECM_SENTINEL_LOGGER_URL", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Errorf("failed to serve: %v", err)
			return
		}
	}()
	defer grpcServer.Stop()

	message := "Test log message"

	ErrorLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[ERROR]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	FatalLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[FATAL]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	WarnLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[WARN]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	InfoLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[INFO]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	AuditLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[AUDIT]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	DebugLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[DEBUG]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}

	TraceLn(&cid, message)
	if !strings.HasPrefix(server.ReceivedMessage, "[TRACE]") && !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}
}

func TestLogServiceSendLog(t *testing.T) {
	// Define a mock LogRequest
	mockRequest := &generated.LogRequest{
		Message: "Test log message",
	}

	// Start a mock gRPC server
	server := &MockLogServiceServer{}
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatalf("error creating log server: %v", err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(lis)

	_ = os.Setenv("VSECM_SENTINEL_LOGGER_URL", lis.Addr().String())

	grpcServer := grpc.NewServer()
	generated.RegisterLogServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Errorf("failed to serve: %v", err)
			return
		}
	}()
	defer grpcServer.Stop()

	// Create a gRPC client
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(conn)
	client := generated.NewLogServiceClient(conn)

	// Call the SendLog method of the LogServiceServer
	response, err := client.SendLog(context.Background(), mockRequest)
	if err != nil {
		t.Fatalf("SendLog failed: %v", err)
	}

	// Check response
	if response == nil {
		t.Error("Response should not be nil")
	}

	if response != nil && response.Message != "" {
		t.Error("Response message should be empty")
	}
	if server.ReceivedMessage != mockRequest.Message {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, mockRequest.Message)
	}
}
