package user

import (
	pb "belajarGrpcGo/proto"
	"context"
	"errors"
	"log"
)

type user struct {
	id    int32
	name  string
	email string
}

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetSingleUser(ctx context.Context, in *pb.IdUser) (foundUser *pb.UserData,err error) {
	log.Printf("Received: %v", in.Id)

	userData := []user{
		{
			name:  "damar",
			id:    1,
			email: "damar@gmail.com",
		},
		{
			name:  "budi",
			id:    2,
			email: "budi@gmail.com",
		},
	}

	for _, v := range userData {
		if in.Id == v.id {
			return &pb.UserData{
				Name:  v.name,
				Id:    v.id,
				Email: v.email,
			}, nil
		}
	}
	return &pb.UserData{}, errors.New("id not found")

}
