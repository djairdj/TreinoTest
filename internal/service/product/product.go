package product

import (
	"TreinoTest/internal/repository/product"
	"TreinoTest/pkg/pb"
	"context"
)

type Product struct {
	productRepository product.Repository
}

func NewProductService(repository product.Repository) pb.ProductServiceServer {
	return &Product{productRepository: repository}
}

func (p *Product) Create(ctx context.Context, request *pb.CreateRequest) (*pb.CreateResponse, error) {
	prod, err := p.productRepository.Create(ctx, request.Name)
	if err != nil {
		return nil, err
	}
	response := &pb.CreateResponse{
		Id:    prod.ID,
		Name:  prod.Name,
		Votes: int32(prod.Votes),
	}
	return response, nil
}

func (p *Product) List(ctx context.Context, request *pb.ListRequest) (*pb.ListResponse, error) {
	lista, err := p.productRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	response := pb.ListResponse{
		Products: []*pb.Product{},
	}
	for _, v := range lista {
		res := pb.Product{
			Id:    v.ID,
			Name:  v.Name,
			Votes: int32(v.Votes),
		}
		response.Products = append(response.Products, &res)
	}
	return &response, err
}

func (p *Product) GetOne(ctx context.Context, request *pb.GetOneRequest) (*pb.GetOneResponse, error) {
	productGot, err := p.productRepository.GetOne(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	pbProduct := pb.Product{
		Id:    productGot.ID,
		Name:  productGot.Name,
		Votes: int32(productGot.Votes),
	}
	res := pb.GetOneResponse{
		Product: &pbProduct,
	}
	return &res, nil
}
