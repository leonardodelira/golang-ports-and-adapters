package productusecase

import (
	"go/cleanarch/core/domain"
	"go/cleanarch/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.Product], error) {
	products, err := usecase.repository.Fetch(paginationRequest)
	if err != nil {
		return nil, err
	}
	return products, nil
}