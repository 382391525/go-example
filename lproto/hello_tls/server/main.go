package main

import (
	"fmt"
	pb "goDemo/lproto/proto/hello"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"
)

const (
	grpcService = ":50052"
)

type helloService struct {
	pb.UnimplementedHelloServer
}

var HelloService = helloService{}

func (h *helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", grpcService)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	//tls 认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/secret.pem", "../../keys/secret.csr")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}
	//实例化grpc server 并开启tls认证
	s := grpc.NewServer(grpc.Creds(creds))
	//注册HelloService
	pb.RegisterHelloServer(s, &HelloService)
	fmt.Println("Listen on: " + grpcService + " tls")
	s.Serve(listen)
}
