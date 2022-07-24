package product

import (
	"TreinoTest/internal/entity"
	"context"
)

type Repository interface {
	Create(ctx context.Context, name string) (*entity.Product, error)
	List(ctx context.Context) ([]entity.Product, error)
	GetOne(ctx context.Context, id string) (*entity.Product, error)
}
