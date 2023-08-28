package main

import (
	"context"
	"fmt"
	h "goapplication/demo3/customer"
	"log"

	"google.golang.org/grpc"
)

func main() {
	con,err := grpc.Dial(":2003",grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed ", err)
	}
	defer con.Close()
	client := h.NewCustomerServiceClient(con)

	customer := &h.CustomerRequest{
		CustomerName: "Naveena",
		AccountId: 123,
		AccountType: "Savings",
	}

	res, err := client.AddCustomer(context.Background(), customer)
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Add: ",res.Result)

	custres, err := client.GetCustomer(context.Background(), &h.Empty{})
	if err != nil {
		log.Fatal("failed ", err)
	}
	fmt.Println("Response of Get: ",custres.Cust)
}