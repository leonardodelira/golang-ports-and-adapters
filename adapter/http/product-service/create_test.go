package productservice_test

import (
	"encoding/json"
	productservice "go/cleanarch/adapter/http/product-service"
	postgresmocks "go/cleanarch/adapter/postgres/mocks"
	"go/cleanarch/core/domain"
	productusecase "go/cleanarch/core/domain/usecase/product-usecase"
	"go/cleanarch/core/dto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func setupCreate() (dto.CreateProductRequest, domain.Product) {
	fakeProductRequest := dto.CreateProductRequest{}
	fakeProduct := domain.Product{}
	faker.FakeData(&fakeProductRequest)
	faker.FakeData(&fakeProduct)

	return fakeProductRequest, fakeProduct
}

func TestCreate(t *testing.T) {
	fakeProductRequest, fakeProduct := setupCreate()

	mockRepo := new(postgresmocks.MockRepository);
	mockRepo.On("Create").Return(&fakeProduct, nil)

	usecase := productusecase.New(mockRepo);
	service := productservice.New(usecase);

	payload, _ := json.Marshal(fakeProductRequest)
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader(string(payload)))
	request.Header.Set("Content-Type", "application/json")
	service.Create(writer, request)

	res := writer.Result()
	defer res.Body.Close()
	
	require.Equal(t, http.StatusOK, res.StatusCode)
}

func TestCreate_JsonErrorFormater(t *testing.T) {
	_, fakeProduct := setupCreate()

	mockRepo := new(postgresmocks.MockRepository);
	mockRepo.On("Create").Return(&fakeProduct, nil)

	usecase := productusecase.New(mockRepo);
	service := productservice.New(usecase);

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/product", strings.NewReader("{"))
	r.Header.Set("Content-Type", "application/json")
	service.Create(w, r)

	res := w.Result()
	defer res.Body.Close()

	require.Equal(t, http.StatusInternalServerError, res.StatusCode)
}