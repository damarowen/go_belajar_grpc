package user

import (
	pb "belajarGrpcGo/protos"
	"context"
	"log"

	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (s *UserService) GetUsers(context.Context, *pb.Empty) (*pb.Users, error) {

	users := &pb.Users{
		Users: []*pb.ResUser{
			{
				Name:  "damar",
				Id:    1,
				Email: "damar@gmail.com",
			},
			{
				Name:  "budi",
				Id:    2,
				Email: "budi@gmail.com",
			},
		},
	}

	return users, nil

}

func (s *UserService) GetSingleUser(ctx context.Context, in *pb.ReqIdUser) (foundUser *pb.ResUser, err error) {
	log.Printf("Received: %v", in.Id)

	userData := []*pb.ResUser{
		{
			Id:    1,
			Name:  "damar",
			Email: "damar@gmail.com",
		},
		{
			Id:    2,
			Name:  "budi",
			Email: "budi@gmail.com",
		},
	}

	for _, v := range userData {
		if in.Id == v.Id {
			return &pb.ResUser{
				Name:  v.Name,
				Id:    v.Id,
				Email: v.Email,
			}, nil
		}
	}
	return nil, status.Error(codes.NotFound, "id was not found")

}

func (s *UserService) CreateSingleUser(ctx context.Context, data *pb.ReqCreateUser) (*pb.Users, error) {

	var max int32

	u := &pb.Users{
		Users: []*pb.ResUser{
			{
				Id:    1,
				Name:  "damar",
				Email: "damar@gmail.com",
			},
			{
				Id:    2,
				Name:  "budi",
				Email: "budi@gmail.com",
			},
		},
	}

	for _, v := range u.Users {
		if v.Id > max {
			max = v.Id
		}
	}

	newUser := pb.ResUser{
		Id:    max + 1,
		Name:  data.GetName(),
		Email: data.GetEmail(),
	}

	u.Users = append(u.Users, &newUser)

	return u, nil

}

func (s *UserService) EditSingleUser(ctx context.Context, data *pb.ReqUser) (*pb.Users, error) {

	u := &pb.Users{
		Users: []*pb.ResUser{
			{
				Id:    1,
				Name:  "damar",
				Email: "damar@gmail.com",
			},
			{
				Id:    2,
				Name:  "budi",
				Email: "budi@gmail.com",
			},
		},
	}

	for _, v := range u.Users {
		if v.Id == data.Id {
			v.Name = data.GetName()
			v.Email = data.GetEmail()
			return u, nil
		}
	}

	return nil, status.Error(codes.NotFound, "id was not found")

}

func (s *UserService) DeleteSingleUser(ctx context.Context, data *pb.ReqIdUser) (*pb.Users, error) {

	var index int

	u := &pb.Users{
		Users: []*pb.ResUser{
			{
				Id:    1,
				Name:  "damar",
				Email: "damar@gmail.com",
			},
			{
				Id:    2,
				Name:  "budi",
				Email: "budi@gmail.com",
			},
		},
	}

	for i, v := range u.Users {
		if v.Id == data.Id {
			index = i
			u.Users = slices.Delete(u.Users, index, index+1)
			return u, nil
		}
	}

	return nil, status.Error(codes.NotFound, "id was not found")

}
