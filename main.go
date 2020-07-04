package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpcclient/services"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 1000})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prodRes)
}
