package di

import (
	productservice "go/cleanarch/adapter/http/product-service"
	"go/cleanarch/adapter/postgres"
	productrepository "go/cleanarch/adapter/postgres/product-repository"
	"go/cleanarch/core/domain"
	productusecase "go/cleanarch/core/domain/usecase/product-usecase"
)

func ConfigProductDI(conn postgres.PoolInterface) domain.ProductService {
	productRepository := productrepository.New(conn)
	productUsecase := productusecase.New(productRepository)
	productService := productservice.New(productUsecase)

	return productService
}