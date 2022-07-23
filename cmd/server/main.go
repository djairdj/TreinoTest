package main

import (
	repositoryproduct "TreinoTest/internal/repository/product"
	serviceproduct "TreinoTest/internal/service/product"
	"TreinoTest/pkg/pb"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	uri := os.Getenv("MONGO_URI")
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	mongodb := client.Database("store")
	repository := repositoryproduct.NewMongoRepository(mongodb)

	server := grpc.NewServer()
	pb.RegisterProductServiceServer(
		server,
		serviceproduct.NewProductService(repository),
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
