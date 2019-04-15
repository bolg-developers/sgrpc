package main

import (
	"log"
	"net"

	"github.com/bolg-developers/sgrpc/server/hello"
	"google.golang.org/grpc"
)

func main() {
	port, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("サーバーが起動しました！")

	server := grpc.NewServer()

	// サービスを登録
	svc := &hello.Service{}
	hello.RegisterHelloServiceServer(server, svc)

	server.Serve(port)
}
