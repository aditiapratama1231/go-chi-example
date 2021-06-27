package product

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
)

type ProductInterface interface {
	GetProducts(ctx context.Context) ([]*entity.Product, error)
	GetSingleProduct(ctx context.Context) (*entity.Product, error)
}
