package api

import (
	"net/http"

	"github.com/aditiapratama1231/graphql-example/internal/brand"
	"github.com/aditiapratama1231/graphql-example/internal/errors"
	"github.com/labstack/echo/v4"
)

type BrandHandler struct {
	brand brand.BrandInterface
}

func NewBrandHandler(usecase brand.BrandInterface) BrandHandler {
	return BrandHandler{
		brand: usecase,
	}
}

func (h BrandHandler) GetBrands(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.brand.GetBrands(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServerError(""))
	}

	return c.JSON(http.StatusOK, result)
}
