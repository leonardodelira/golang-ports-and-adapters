package productservice

import (
	"encoding/json"
	"net/http"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	products, err := service.usecase.Fetch()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(products)
}