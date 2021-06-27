package repository

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
	"github.com/jinzhu/gorm"
)

type ProductRepositoryInterface interface {
	GetProducts(ctx context.Context) ([]*entity.Product, error)
	GetSingleProduct(ctx context.Context) (*entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (pr ProductRepository) GetProducts(ctx context.Context) ([]*entity.Product, error) {
	var products []*entity.Product

	pr.db.Find(&products)

	return products, nil
}

func (pr ProductRepository) GetSingleProduct(ctx context.Context) (*entity.Product, error) {
	product := &entity.Product{}

	pr.db.First(product)

	return product, nil
}
