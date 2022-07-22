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
	server := grpc.NewServer()
	pb.RegisterProductServiceServer(
		server,
		product.NewProductService(),
	)
	port := ":50051"
	reflection.Register(server)
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
