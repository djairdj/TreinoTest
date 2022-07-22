package product

import (
	"TreinoTest/pkg/pb"
	"context"
	"log"
)

type Product struct {
}

func NewProductService() pb.ProductServiceServer {
	return &Product{}
}

func (*Product) CreateProduct(ctx context.Context, request *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	response := &pb.CreateProductResponse{Name: request.Name}
	log.Println("Mensagem recebida", request.Name)
	return response, nil
}
