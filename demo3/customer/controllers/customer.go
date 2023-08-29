package rpcService

import (
	"context"
	"fmt"
	pb "goapplication/demo3/customer"
	"goapplication/demo3/interfaces"
	"goapplication/demo3/models"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)


type RPCServer struct {
	mu        sync.Mutex
	pb.UnimplementedCustomerServiceServer

}

var(
    CustomerService  interfaces.Icustomer
	Mcoll *mongo.Collection
)

func (s *RPCServer) AddCustomer(ctx context.Context, req *pb.CustomerRequired) (*pb.CustomerResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fmt.Println(req)
	cust := &models.Customer{
		Customer_ID: req.CustomerId,
		Bank_ID:     req.BankId,
		Password:    req.Password,
		Name:        req.Name,
		Account_ID:  req.AccountId,
		Balance:     req.Balance,
	}
	fmt.Println(cust)
	res,err := CustomerService.CreateCustomer(cust)
	if err!=nil{
		return nil,err
	}
	fmt.Println(res)
	return &pb.CustomerResponse{
		Result: "Success",
	}, nil
}

func (s *RPCServer) GetCustomer(ctx context.Context, req *pb.CustID) (*pb.CustomerRequired, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	res, err := CustomerService.GetCustomerById(req.CustomerId)
	if err!=nil{
		return nil,err
	}
	customer := &pb.CustomerRequired{
		CustomerId: res.Customer_ID,
		BankId:     res.Bank_ID,
		Password:   res.Password,
		Name:       res.Name,
		AccountId:  res.Account_ID,
		Balance:    res.Balance,
	}
	return customer, nil
}

func (s *RPCServer) UpdateCustomer(ctx context.Context, req *pb.UpdateReq) (*pb.CustomerResponse, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	res,err := CustomerService.UpdateCustomerById(req.Id, &models.UpdateModel{Topic: req.Topic, FinalValue: req.Newvalue})
	if err!=nil{
		return nil,err
	}
	fmt.Println(res)
	return &pb.CustomerResponse{
		Result: "Success",
	}, nil
}

func (s *RPCServer) DeleteCustomer(ctx context.Context, req *pb.CustID) (*pb.CustomerResponse, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	res,err := CustomerService.DeleteCustomerById(req.CustomerId)
	if err!=nil{
		return nil,err
	}
	fmt.Println(res)
	return &pb.CustomerResponse{
		Result: "Success",
	}, nil
}

