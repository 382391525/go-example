package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "learn-golang/lproto/proto/hello"
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
	//实例化grpc server
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &HelloService)
	grpclog.Infoln("Listen on " + grpcService)
	s.Serve(listen)
}
