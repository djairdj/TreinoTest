package product

import (
	"TreinoTest/entity"
	"context"
)

type Repository interface {
	Create(ctx context.Context, name string) (*entity.Product, error)
}
