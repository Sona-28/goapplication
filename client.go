package main

import (
	"context"
	"fmt"
	h "goapplication/helloworld"
	"log"

	"google.golang.org/grpc"
)

func main(){
	conn,err := grpc.Dial("localhost:2023",grpc.WithInsecure())
	if err!=nil{
		fmt.Println("Failed to connect: ",err)
	}
	defer conn.Close()
	client := h.NewGreeterClient(conn)

	name := "Naveena"
	response, err := client.SayHello(context.Background(),&h.HelloRequest{Name: name})
	if err!=nil{
		log.Fatal("Failed to call: ",err)
	}
	fmt.Println("Response: ",response.Message)
}