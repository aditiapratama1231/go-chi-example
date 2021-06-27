package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/aditiapratama1231/graphql-example/pkg/database"

	brand_api "github.com/aditiapratama1231/graphql-example/internal/brand/api"
	prod_api "github.com/aditiapratama1231/graphql-example/internal/product/api"

	brand_usecase "github.com/aditiapratama1231/graphql-example/internal/brand/use_case"
	prod_usecase "github.com/aditiapratama1231/graphql-example/internal/product/use_case"

	brand_repo "github.com/aditiapratama1231/graphql-example/internal/brand/repository"
	prod_repo "github.com/aditiapratama1231/graphql-example/internal/product/repository"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// initial handler
	db := database.DBInit()
	productHandler := productHandlerInit(db)
	brandHandler := brandHandlerInit(db)

	// product
	e.GET("/products", productHandler.GetProducts)
	e.GET("/product", productHandler.GetProduct)

	// brand
	e.GET("/brands", brandHandler.GetBrands)

	e.Logger.Fatal(e.Start(":" + port))
}

func productHandlerInit(db *gorm.DB) prod_api.ProductHandler {
	productRepo := prod_repo.NewProductRepository(db)
	productUsecase := prod_usecase.NewProductResolver(productRepo)
	return prod_api.NewProductHandler(productUsecase)
}

func brandHandlerInit(db *gorm.DB) brand_api.BrandHandler {
	brandRepo := brand_repo.NewBrandRepository(db)
	brandUsecase := brand_usecase.NewProductResolver(brandRepo)
	return brand_api.NewBrandHandler(brandUsecase)
}
