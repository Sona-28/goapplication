package main

import (
	"context"
	"fmt"
	"log"

	pb "goapplication/demo2/task"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:2002", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed1 %v", err)
	}
	defer conn.Close()
	client2 := pb.NewTaskServiceClient(conn)

	task := &pb.Task{
		Title: "Buy groceries",
	}

	addres, err := client2.AddTask(context.Background(), task)
	if err != nil {
		log.Fatal("failed2 %v", err)
	}

	fmt.Printf("Response of add :%s\n", addres.Id)

	taskres, err := client2.GetTasks(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatal("failed2 %v", err)
	}
	fmt.Printf("Response of get :")
	for _, task := range taskres.Tasks {
		fmt.Printf("ID:%s,Title:%s,Completed:%v\n", task.Id, task.Title, task.Completed)
	}
}
