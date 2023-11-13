package main

import (
	"context"
	"log"

	pb "github.com/CtrlShiftAlt/micro-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect:%v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	req := &pb.HelloRequest{Name: "John"}
	reply, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call SayHello:%v", err)
	}
	log.Printf("Reply from SayHello: %s", reply.Message)
}
