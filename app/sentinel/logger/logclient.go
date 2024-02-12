package logger

import (
	"context"
	pb "github.com/vmware-tanzu/secrets-manager/app/sentinel/logger/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func SendLogMessage(a ...any) {
	conn, err := grpc.Dial(
		LOGGER_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Printf("Logger.SendLogMessage could not connect to server: %v\n", err)
		return
	}
	defer conn.Close()
	c := pb.NewLogServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.SendLog(ctx, &pb.LogRequest{Message: LogTextBuilder(a...)})
	if err != nil {
		log.Printf("Logger.SendLogMessage could not send message: %v\n", err)
		return
	}
}
