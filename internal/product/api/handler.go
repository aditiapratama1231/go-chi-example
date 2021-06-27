package api

import (
	"net/http"

	"github.com/aditiapratama1231/graphql-example/internal/errors"
	"github.com/aditiapratama1231/graphql-example/internal/product"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	product product.ProductInterface
}

func NewProductHandler(usecase product.ProductInterface) ProductHandler {
	return ProductHandler{
		product: usecase,
	}
}

func (h ProductHandler) GetProducts(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.product.GetProducts(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServerError(""))
	}

	return c.JSON(http.StatusOK, result)
}

func (h ProductHandler) GetProduct(c echo.Context) error {
	ctx := c.Request().Context()
	result, err := h.product.GetSingleProduct(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, errors.InternalServerError(""))
	}

	return c.JSON(http.StatusOK, result)
}
