package main

import (
	pb "belajarGrpcGo/proto"
	"belajarGrpcGo/server/user"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

//di video tutorial ada yang bikin kaya gini
// trus di passing sebelahan sama grpcServer pas register
type UserServiceServer struct {
	pb.UserServiceServer
}

func main() {

	//initiate connection
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create new gRPC server
	grpcServer := grpc.NewServer()

	// create new instance of User service server
	userServ := user.NewUserService()

	// register user service server to grpc
	pb.RegisterUserServiceServer(grpcServer, userServ)

	log.Printf("server listening at %v", lis.Addr())

	//listening grpc server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
