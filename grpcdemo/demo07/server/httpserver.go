package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"grpcstudy/demo07/server/services"
	"log"
	"net"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	gwmux := runtime.NewServeMux()
	defer cancel()
	/*	cert, _ := tls.LoadX509KeyPair("cert/client.pem","cert/client.key")
		certPool := x509.NewCertPool()
		ca, _ := ioutil.ReadFile("cert/ca.pem")
		certPool.AppendCertsFromPEM(ca)

		creds := credentials.NewTLS(&tls.Config{
			Certificates: []tls.Certificate{cert},//服务端证书
			ServerName: "localhost",
			RootCAs: certPool,
		})*/
	opt := []grpc.DialOption{grpc.WithInsecure()}
	//opt := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	//8099是对应的grpc的监听端口，必须与grpc开放的端口一致,所以启动前务必启动grpc服务
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8099", opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			fmt.Println("启动http server faield", err)
		}
	}()

	//rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp", ":8096")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(lis)

}
