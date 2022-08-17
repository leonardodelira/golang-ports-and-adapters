package productservice

import (
	"encoding/json"
	"go/cleanarch/core/dto"
	"net/http"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	pagination, err := dto.FromValuePaginationRequestParams(request)
	
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	products, err := service.usecase.Fetch(pagination)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}