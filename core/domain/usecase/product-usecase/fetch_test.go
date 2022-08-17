package productusecase_test

import (
	"errors"
	postgresmocks "go/cleanarch/adapter/postgres/mocks"
	"go/cleanarch/core/domain"
	productusecase "go/cleanarch/core/domain/usecase/product-usecase"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	var product = domain.Product{}
	faker.FakeData(&product)

	pagination := domain.Pagination[[]domain.Product]{
		Items: []domain.Product{product},
		Total: 0,
	};

	mockRepo := new(postgresmocks.MockRepository)
	mockRepo.On("Fetch").Return(&pagination, nil)

	sut := productusecase.New(mockRepo)
	response, err := sut.Fetch()

	require.Nil(t, err)
	require.NotNil(t, response.Items)
	require.Equal(t, int(1), len(response.Items))
	require.Equal(t, int32(0), response.Total)
}

func TestFetch_Error(t *testing.T) {
	var product = domain.Product{}
	faker.FakeData(&product)

	pagination := domain.Pagination[[]domain.Product]{};

	mockRepo := new(postgresmocks.MockRepository)
	mockRepo.On("Fetch").Return(&pagination, errors.New("SOME ERROR"))

	sut := productusecase.New(mockRepo)
	response, err := sut.Fetch()

	require.NotNil(t, err)
	require.Nil(t, response)
}