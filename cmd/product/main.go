package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jinzhu/gorm"

	"github.com/aditiapratama1231/graphql-example/graph"
	"github.com/aditiapratama1231/graphql-example/graph/dataloader"
	"github.com/aditiapratama1231/graphql-example/graph/generated"
	"github.com/aditiapratama1231/graphql-example/pkg/database"

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

	db := database.DBInit()
	dl := dataloader.NewRetriever()

	var resolver = &graph.Resolver{
		ProductResolver: productResolverInit(db),
		BrandResolver:   brandResolverInit(db),
		DataLoaders:     dl,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func productResolverInit(db *gorm.DB) prod_usecase.Product {
	productRepo := prod_repo.NewProductRepository(db)
	return prod_usecase.Product{
		ProductRepository: productRepo,
	}
}

func brandResolverInit(db *gorm.DB) brand_usecase.Brand {
	brandRepo := brand_repo.NewBrandRepository(db)
	return brand_usecase.Brand{
		BrandRepository: brandRepo,
	}
}
