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
	stream, err := userClient.GetUserInfoByClientStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 3; i ++ {
		userIdList := make([]int32, 0)
		for j := 1; j <= 10; j ++ {
			userIdList = append(userIdList, int32(1000 * i + j))
		}
		req := &services.UserRequest{UserIds:userIdList}

		err := stream.Send(req)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Users)
}
