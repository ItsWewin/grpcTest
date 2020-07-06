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
	//creds, err := credentials.NewClientTLSFromFile("certificate/backend.cert", "wewin")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//certPool := x509.NewCertPool()
	//ca, err := ioutil.ReadFile("cert/ca.pem")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//certPool.AppendCertsFromPEM(ca)
	//
	//creds := credentials.NewTLS(&tls.Config{
	//	Certificates: []tls.Certificate{cert},
	//	ServerName:   "localhost",
	//	RootCAs: 	  certPool,
	//})

	creds, err := helper.GetClientCreds()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	prodClient := services.NewProdServiceClient(conn)
	//prodRes, err := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 1000})
	//if err != nil {
	//	log.Fatal(err)
	//}

	prodRes, err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size:10})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(prodRes)
}
