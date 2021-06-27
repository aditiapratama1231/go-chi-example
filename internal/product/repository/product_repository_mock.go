package repository

import (
	"context"
	"errors"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repo *ProductRepositoryMock) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	products := []*entity.Product{}
	return products, nil
}

func (repo *ProductRepositoryMock) GetSingleProduct(ctx context.Context) (*entity.Product, error) {
	arguments := repo.Mock.Called(ctx)

	if arguments.Get(0) == nil {
		return nil, errors.New("Product not found")
	}

	product := arguments.Get(0).(*entity.Product)
	return product, nil
}
