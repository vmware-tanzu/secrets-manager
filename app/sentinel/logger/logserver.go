package logger

import (
	"context"
	"fmt"
	pb "github.com/vmware-tanzu/secrets-manager/app/sentinel/logger/generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedLogServiceServer
}

func (s *server) SendLog(ctx context.Context, in *pb.LogRequest) (*pb.LogResponse, error) {
	fmt.Printf("%s", in.Message)
	return &pb.LogResponse{}, nil
}

func CreateLogServer() {
	lis, err := net.Listen("tcp", SentinelLoggerUrl())
	if err != nil {
		log.Printf("Logger.CreateLogServer error creating log server: %v\n", err)
		return
	}
	s := grpc.NewServer()

	pb.RegisterLogServiceServer(s, &server{})

	log.Printf("Logger.CreateLogServer listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("Logger.CreateLogServer failed to serve log server: %v\n", err)
		return
	}
}
