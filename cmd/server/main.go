package main

import (
	"TreinoTest/pkg/pb"
	"TreinoTest/pkg/srv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	servidor := grpc.NewServer()
	pb.RegisterProductsServiceServer(servidor, &srv.ProductServer{})
	port := ":5001"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpc_Error := servidor.Serve(listener)
	if grpc_Error != nil {
		log.Fatal(grpc_Error)
	}
}
