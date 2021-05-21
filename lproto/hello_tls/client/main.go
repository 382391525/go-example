package main

import (
	"context"
	"fmt"
	pb "goDemo/lproto/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	grpcService = ":50052"
)

func main()  {
	// TLS连接  记得把server name改成你写的服务器地址
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "www.fuzhifei.com")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(grpcService,grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)
	req := &pb.HelloRequest{Name: "grpc"}
	resp, err := c.SayHello(context.Background(),req)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(resp.Message)
}