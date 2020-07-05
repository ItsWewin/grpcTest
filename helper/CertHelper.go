package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

func GetClientCreds() (credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("cert/ca.pem")
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs: 	  certPool,
	})

	return creds, nil
}