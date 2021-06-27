package use_case

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/brand/repository"
	"github.com/aditiapratama1231/graphql-example/internal/entity"
)

// brand service
type Brand struct {
	BrandRepository repository.BrandRepositoryInterface
}

func NewProductResolver(productRepo repository.BrandRepositoryInterface) Brand {
	return Brand{
		BrandRepository: productRepo,
	}
}

func (b Brand) GetBrands(ctx context.Context) ([]*entity.Brand, error) {
	result, err := b.BrandRepository.GetBrands(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
