package main

import (
	"TreinoTest/pkg/pb"
	"TreinoTest/pkg/srv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	servidor := grpc.NewServer()
	pb.RegisterProductsServiceServer(
		servidor,
		&srv.ProductServer{},
	)
	port := ":5001"
	reflection.Register(servidor)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpc_Error := servidor.Serve(listener)
	if grpc_Error != nil {
		log.Fatal(grpc_Error)
	}
}
