package repository

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
	"github.com/jinzhu/gorm"
)

type BrandRepositoryInterface interface {
	GetBrands(ctx context.Context) ([]*entity.Brand, error)
}

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) BrandRepositoryInterface {
	return &BrandRepository{
		db: db,
	}
}

func (br BrandRepository) GetBrands(ctx context.Context) ([]*entity.Brand, error) {
	var brands []*entity.Brand

	br.db.Preload("Products").Find(&brands)

	return brands, nil
}
