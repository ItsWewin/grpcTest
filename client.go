package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
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
	userIdList := []int32{1, 2, 3, 4, 5, 6}
	resp, err := userClient.GetUsersInfo(context.Background(), &services.UserRequest{UserIds: userIdList})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Users)
}
