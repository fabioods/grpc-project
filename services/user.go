package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/fabioods/fullcyle-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }
//AddUserVerbose(ctx context.Context, in *User, opts ...grpc.CallOption) (UserService_AddUserVerboseClient, error)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {

	fmt.Println("AddUser ", req.Name)

	return &pb.User{
		Id:    "1",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (u *UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Inserting",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	time.Sleep(time.Second * 3)
	return nil
}

func (u *UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{User: users})
		}
		if err != nil {
			log.Fatalf("failed to receive msg: %v", err)
		}
		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.GetName(),
			Email: req.GetEmail(),
		})
		fmt.Println("AddUsers ", req.GetName())
	}
}

func (u *UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("failed to receive msg: %v", err)
		}
		err = stream.Send(&pb.UserResultStream{
			Status: "Inserting",
			User:   req,
		})
		if err != nil {
			log.Fatalf("failed to send msg: %v", err)
		}
		fmt.Println("AddUserStreamBot ", req.GetName())
	}
}
