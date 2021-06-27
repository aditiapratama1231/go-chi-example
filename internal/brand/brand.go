package brand

import (
	"context"

	"github.com/aditiapratama1231/graphql-example/internal/entity"
)

type BrandInterface interface {
	GetBrands(ctx context.Context) ([]*entity.Brand, error)
}
