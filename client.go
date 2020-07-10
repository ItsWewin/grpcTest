package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"log"
	"mayihahah.com/grpc/helper"
	"mayihahah.com/grpc/services"
	"time"
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

	orderClient := services.NewOrderServiceClient(conn)
	t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	resp, err := orderClient.NewOrder(context.Background(), &services.OrderRequest{OrderMain: &services.OrderMain{
		OrderId:     12345677,
		OrderNo:    "orderNO-12345677",
		UserId:     1,
		OrderMoney: 1234.567,
		OrderTime:  &t,
	}} )
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(resp)
}
