package productusecase_test

import (
	"fmt"
	"go/cleanarch/core/domain"
	"go/cleanarch/core/dto"
	"testing"

	productusecase "go/cleanarch/core/domain/usecase/product-usecase"

	postgresmocks "go/cleanarch/adapter/postgres/mocks"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	fakeDBProduct := domain.Product{}
	faker.FakeData(&fakeRequestProduct)
	faker.FakeData(&fakeDBProduct)
	
	mockRepo := new(postgresmocks.MockRepository)
	mockRepo.On("Create").Return(&fakeDBProduct, nil)

	sut := productusecase.New(mockRepo)
	product, err := sut.Create(&fakeRequestProduct)

	require.Nil(t, err)
	require.NotEmpty(t, product.ID)
	require.Equal(t, product.Name, fakeDBProduct.Name)
	require.Equal(t, product.Price, fakeDBProduct.Price)
	require.Equal(t, product.Description, fakeDBProduct.Description)
}

func TestCreate_Error(t *testing.T) {
	fakeRequestProduct := dto.CreateProductRequest{}
	faker.FakeData(&fakeRequestProduct)

	mockRepo := new(postgresmocks.MockRepository)
	mockRepo.On("Create").Return(&domain.Product{}, fmt.Errorf("SOME ERROR"))

	sut := productusecase.New(mockRepo)
	_, err := sut.Create(&fakeRequestProduct)

	require.NotNil(t, err)
}