package services

import (
	"context"
	"fmt"

	"github.com/fabioods/fullcyle-grpc/pb"
)

// type UserServiceServer interface {
// 	AddUser(context.Context, *User) (*User, error)
// 	mustEmbedUnimplementedUserServiceServer()
// }

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
