package main

import (
	"context"
	"fmt"
	"log"

	"github.com/fabioods/fullcyle-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	AddUser(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Fabio",
		Email: "fah_ds@live.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}
	fmt.Println("AddUser ", res)
}
