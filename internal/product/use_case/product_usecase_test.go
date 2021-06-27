package use_case

import (
	"context"
	"testing"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
	"github.com/aditiapratama1231/graphql-example/internal/product/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Setup
var (
	productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
	productUseCase    = Product{ProductRepository: productRepository}
	products          = []*entity.Product{
		&entity.Product{
			ID:    1,
			Name:  "Logitech",
			Stock: 10,
		},
		&entity.Product{
			ID:    2,
			Name:  "Razer",
			Stock: 3,
		},
	}
)

func TestProductUseCase_GetProducts(t *testing.T) {
	ctx := context.Background()

	productRepository.Mock.On("GetProducts", ctx).Return(products)
	result, err := productUseCase.GetProducts(ctx)
	if err != nil {
		t.Errorf("something wrong with GetProducts")
	}
	assert.NotNil(t, result)
}

func TestProductUseCase_GetSingleProduct(t *testing.T) {
	ctx := context.Background()

	productRepository.Mock.On("GetSingleProduct", ctx).Return(products[0])
	result, err := productUseCase.GetSingleProduct(ctx)
	if err != nil {
		t.Errorf("something wrong with GetSingleProduct")
	}
	assert.NotNil(t, result)
}
