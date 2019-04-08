package main

import (
	"github.com/bolg-developers/sgrpc/server/hello"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	port, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	// サービスを登録
	svc := &hello.Service{}
	hello.RegisterHelloServiceServer(server, svc)

	server.Serve(port)
}
