package productusecase

import (
	"go/cleanarch/core/domain"
)

func (usecase usecase) Fetch() (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repository.Fetch()
	if err != nil {
		return nil, err
	}
	return products, nil
}