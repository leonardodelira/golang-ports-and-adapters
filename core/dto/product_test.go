package dto_test

import (
	"encoding/json"
	"go/cleanarch/core/dto"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

func TestFromJSONCreateProductRequest(t *testing.T) {
	fakeItem := dto.CreateProductRequest{}
	faker.FakeData(&fakeItem)

	json, err := json.Marshal(fakeItem)
	require.Nil(t, err)

	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader(string(json)))

	require.Nil(t, err)
	require.NotNil(t, itemRequest)
	require.Equal(t, fakeItem.Name, itemRequest.Name)
	require.Equal(t, fakeItem.Description, itemRequest.Description)
	require.Equal(t, fakeItem.Price, itemRequest.Price)
}

func TestFromJSONCreateProductRequest_JSONDecodeError(t *testing.T) {
	itemRequest, err := dto.FromJSONCreateProductRequest(strings.NewReader("{"))

	require.NotNil(t, err)
	require.Nil(t, itemRequest)
}