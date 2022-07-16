package main

import (
	"TreinoTest/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	servidor := grpc.NewServer()
	pb.RegisterProductsServiceServer(servidor)
}
