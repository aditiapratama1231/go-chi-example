package use_case

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
	"github.com/aditiapratama1231/graphql-example/internal/product/repository"
)

// product service
type Product struct {
	ProductRepository repository.ProductRepositoryInterface
}

func NewProductResolver(productRepo repository.ProductRepositoryInterface) Product {
	return Product{
		ProductRepository: productRepo,
	}
}

func (p Product) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	result, err := p.ProductRepository.GetProducts(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (p Product) GetSingleProduct(ctx context.Context) (*entity.Product, error) {
	result, err := p.ProductRepository.GetSingleProduct(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
