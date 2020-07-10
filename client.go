package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mayihahah.com/grpc/helper"
	"mayihahah.com/grpc/services"
)

func main() {
	creds, err := helper.GetClientCreds()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	userClient := services.NewUserServiceClient(conn)
	userIdList := []int32{1, 2, 3, 4, 5, 6, 7}
	stream, err := userClient.GetUserInfoByServerStream(context.Background(), &services.UserRequest{UserIds: userIdList})
	if err != nil {
		log.Fatal(err)
	}

	for {
		res, err := stream.Recv()
		switch {
		case err != io.EOF:
			break
		case err != nil:
			log.Fatal(err)
		}

		fmt.Println(res.Users)
	}
}
