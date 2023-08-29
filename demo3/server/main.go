package main

import (
	"context"
	"fmt"
	"goapplication/demo3/config"
	"goapplication/demo3/constants"
	pb "goapplication/demo3/customer"
	rpc "goapplication/demo3/customer/controllers"
	"goapplication/demo3/service"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)


func initApp(mongoClient *mongo.Client){
	rpc.Mcoll = config.GetCollection(mongoClient, constants.Dbname, "customer")
	rpc.CustomerService = service.InitCustomer(rpc.Mcoll, context.Background())
}

func main() {
	mongoClient,err := config.ConnectDataBase()
	defer mongoClient.Disconnect(context.TODO())
	if err!=nil{
		panic(err)
	}
	initApp(mongoClient)
	lis, err := net.Listen("tcp", constants.Port)
	fmt.Println("Server listening on: ", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen:%v", err)
		return
	}
	s := grpc.NewServer()
	// fmt.Println("server")
	pb.RegisterCustomerServiceServer(s,&rpc.RPCServer{})
	// fmt.Println("server1")
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve:%v", err)
	}
	fmt.Println("finish")
	
	
}