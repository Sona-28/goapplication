package main

import (
	"context"
	"fmt"
	"net"
	h "goapplication/helloworld"
	"google.golang.org/grpc"
)
type server struct{
	h.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, rq *h.HelloRequest) (*h.HelloResponse, error) {
	return &h.HelloResponse{
		Message: fmt.Sprintf("Hello %s", rq.Name),
	},nil
}

func main(){
	lis,err := net.Listen("tcp",":2023")
	if err!=nil{
		fmt.Println("Failed to listen: ",err)
		return
	}
	s := grpc.NewServer()
	h.RegisterGreeterServer(s,&server{})

	fmt.Println("Server listening on port: 2023")
	if err := s.Serve(lis); err!=nil{
		fmt.Println("Failed to serve: ",err)
		return
	}
}

