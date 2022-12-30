// Package main implements a client for Greeter service.
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "belajarGrpcGo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	serverPort = ":9000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	//connect to proto
	c := pb.NewUserServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	idUser := 1
	r, err := c.GetSingleUser(ctx, &pb.IdUser{Id: int32(idUser)})
	if err != nil {
		log.Fatalf("could not found user: %v", err)
	}

	fmt.Println(r.Name)
	fmt.Println(r.Email)
	fmt.Println(r.Id)

}

