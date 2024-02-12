package logger

import (
	"context"
	"net"
	"strings"
	"testing"

	pb "github.com/vmware-tanzu/secrets-manager/app/sentinel/logger/generated"
	"google.golang.org/grpc"
)

type MockLogServiceServer struct {
	pb.UnimplementedLogServiceServer
	ReceivedMessage string
}

func (s *MockLogServiceServer) SendLog(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	s.ReceivedMessage = in.Message
	return &pb.LogResponse{}, nil
}

func TestSendLogMessage(t *testing.T) {
	server := &MockLogServiceServer{}
	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, server)

	lis, err := net.Listen("tcp", LOGGER_PORT)
	if err != nil {
		t.Fatalf("error creating log server: %v", err)
	}
	defer lis.Close()

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()
	defer grpcServer.Stop()

	message := "Test log message"
	SendLogMessage(message)

	if !strings.Contains(server.ReceivedMessage, message) {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, message)
	}
}

func TestCreateLogServer(t *testing.T) {
	server := &MockLogServiceServer{}
	lis, err := net.Listen("tcp", LOGGER_PORT)
	if err != nil {
		t.Fatalf("error creating log server: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()
	defer grpcServer.Stop()

	CreateLogServer()
}

func TestLogServiceSendLog(t *testing.T) {
	// Define a mock LogRequest
	mockRequest := &pb.LogRequest{
		Message: "Test log message",
	}

	// Start a mock gRPC server
	server := &MockLogServiceServer{}
	lis, err := net.Listen("tcp", LOGGER_PORT)
	if err != nil {
		t.Fatalf("error creating log server: %v", err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	pb.RegisterLogServiceServer(grpcServer, server)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			t.Fatalf("failed to serve: %v", err)
		}
	}()
	defer grpcServer.Stop()

	// Create a gRPC client
	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Fatalf("failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewLogServiceClient(conn)

	// Call the SendLog method of the LogServiceServer
	response, err := client.SendLog(context.Background(), mockRequest)
	if err != nil {
		t.Fatalf("SendLog failed: %v", err)
	}

	// Check response
	if response == nil {
		t.Error("Response should not be nil")
	}
	if response.Message != "" {
		t.Error("Response message should be empty")
	}
	if server.ReceivedMessage != mockRequest.Message {
		t.Errorf("Received message (%s) does not match sent message (%s)", server.ReceivedMessage, mockRequest.Message)
	}
}
