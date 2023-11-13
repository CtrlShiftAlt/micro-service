package main

import (
	"context"
	"log"
	"net"

	pb "github.com/CtrlShiftAlt/micro-service/service"
	"google.golang.org/grpc"
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	message := "你好" + req.Name + ", Hello, " + req.Name
	return &pb.HelloReply{Message: message}, nil
}

func main() {
	// 监听本地端口
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	// 创建gRPC服务
	server := grpc.NewServer()
	// 注册服务
	pb.RegisterGreeterServer(server, &greeterServer{})

	log.Println("gRPC server is listening on port 50051")

	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
