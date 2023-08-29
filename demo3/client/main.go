package main

import (
	"context"
	"fmt"
	"goapplication/demo3/constants"
	h "goapplication/demo3/customer"
	"log"
	"google.golang.org/grpc"
)

func Create(client h.CustomerServiceClient) {
	customer := &h.CustomerRequired{
		CustomerId: 102,
		BankId:     1001,
		Password:   "naveena",
		Name:       "Naveena",
		AccountId:  789457123,
		Balance:    7000,
	}

	res, err := client.AddCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Add: ", res.Result)
}

func Read(client h.CustomerServiceClient) {
	custres, err := client.GetCustomer(context.Background(), &h.CustID{
		CustomerId: 101,
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Get: ", custres)
}

func Update(client h.CustomerServiceClient) {
	res, err := client.UpdateCustomer(context.Background(), &h.UpdateReq{
		Id:       101,
		Topic:    "name",
		Newvalue: "Harry",
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Update: ", res.Result)
}

func Delete(client h.CustomerServiceClient) {

	res, err := client.DeleteCustomer(context.Background(), &h.CustID{
		CustomerId: 102,
	})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Delete: ", res.Result)
}

func main() {
	con, err := grpc.Dial(constants.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed ", err)
	}
	defer con.Close()
	client := h.NewCustomerServiceClient(con)

	Create(client)

	// Update(client)

	// Read(client)

	// Delete(client)

}
