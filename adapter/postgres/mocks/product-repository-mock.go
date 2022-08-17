package postgresmocks

import (
	"go/cleanarch/core/domain"
	"go/cleanarch/core/dto"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock MockRepository) Create(productRequest *dto.CreateProductRequest) (*domain.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Product), args.Error(1)
}

func (mock MockRepository) Fetch() (*domain.Pagination[[]domain.Product], error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*domain.Pagination[[]domain.Product]), args.Error(1)
}


