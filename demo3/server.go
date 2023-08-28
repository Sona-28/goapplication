package main

import (
	"context"
	"fmt"
	pb "goapplication/demo3/customer"
	"net"
	"strconv"
	"sync"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)



type customerServer struct {
	mu        sync.Mutex
	customers map[int32]*pb.CustomerRequest
	pb.UnimplementedCustomerServiceServer
}

func (s *customerServer) AddCustomer(ctx context.Context, req *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	customerID := generateID()
	req.CustomerId = customerID
	s.customers[req.CustomerId] = req

	return &pb.CustomerResponse{
		Result: "Success",
	}, nil
}

func (s *customerServer) GetCustomer(ctx context.Context, req *pb.Empty) (*pb.CustList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	customers := make([]*pb.CustomerRequest, 0, len(s.customers))
	for _, customer := range s.customers {
		customers = append(customers, customer)
	}
	return &pb.CustList{Cust: customers}, nil
}

func generateID() int32 {
	id:= primitive.NewObjectID()
	id1 := id.String()
	cusid, _:= strconv.ParseInt(id1, 10, 32)
	return int32(cusid)
}

func main() {
	lis, err := net.Listen("tcp", ":2003")
	fmt.Println("Server listening on:2003")
	if err != nil {
		fmt.Printf("Failed to listen:%v", err)
		return
	}
	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, &customerServer{
		customers: make(map[int32]*pb.CustomerRequest),
	})
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve:%v", err)
	}
	
	
}