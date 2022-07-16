package main

import (
	"TreinoTest/internal/service/product"
	"TreinoTest/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	servidor := grpc.NewServer()
	pb.RegisterProductsServiceServer(
		servidor,
		&product.ProductServer{},
	)
	port := ":5001"
	reflection.Register(servidor)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcError := servidor.Serve(listener)
	if grpcError != nil {
		log.Fatal(grpcError)
	}
}
