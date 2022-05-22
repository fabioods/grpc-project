package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

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

	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUserStreamBoth(client)
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

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Fabio",
		Email: "fah_ds@live.com",
	}
	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}
	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive msg: %v", err)
		}
		fmt.Println("Status", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("failed to add user: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive response: %v", err)
	}
	fmt.Println(res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	stream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	reqs := []*pb.User{
		&pb.User{
			Id:    "0",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "1",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "4",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
		&pb.User{
			Id:    "5",
			Name:  "Fabio",
			Email: "fah_ds@live.com",
		},
	}

	wait := make(chan int)

	go func() {
		for _, req := range reqs {
			fmt.Println("Sending req: ", req)
			stream.Send(req)
			time.Sleep(time.Second * 3)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("failed to receive msg: %v", err)
				break
			}
			fmt.Println("User", res)
		}
		close(wait)
	}()

	<-wait
}
