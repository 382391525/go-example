package main

import (
	"context"
	"fmt"
	pb "goDemo/lproto/proto/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	grpcService = ":50052"
)

func main()  {
	conn, err := grpc.Dial(grpcService,grpc.WithInsecure())
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