package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"mayihahah.com/grpcclient/services"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("certificate/backend.cert", "wewin")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
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
