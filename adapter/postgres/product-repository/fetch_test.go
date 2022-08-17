package productrepository_test

import (
	"fmt"
	productrepository "go/cleanarch/adapter/postgres/product-repository"
	"go/cleanarch/core/domain"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func setupFetch() ([]string, domain.Product, pgxmock.PgxPoolIface) {
	cols := []string{"id", "name", "price", "description"}
	
	fakeProductDBResponse := domain.Product{}
	faker.FakeData(&fakeProductDBResponse)

	mock, _ := pgxmock.NewPool()

	return cols, fakeProductDBResponse, mock
}

func TestFetch(t *testing.T) {
	cols, fakeProductDBResponse, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnRows(
		pgxmock.NewRows(cols).AddRow(
			fakeProductDBResponse.ID,
			fakeProductDBResponse.Name,
			fakeProductDBResponse.Price,
			fakeProductDBResponse.Description,
		))

	sut := productrepository.New(mock)
	products, err := sut.Fetch()

	require.Nil(t, err)
	require.NotNil(t, products)

	for _, product := range products.Items {
		require.Equal(t, fakeProductDBResponse.Name, product.Name)
		require.Equal(t, fakeProductDBResponse.Description, product.Description)
		require.Equal(t, fakeProductDBResponse.Price, product.Price)
	}
}

func TestFetch_QueryError(t *testing.T) {
	_, _, mock := setupFetch()
	defer mock.Close()

	mock.ExpectQuery("SELECT (.+) FROM product").WillReturnError(
		fmt.Errorf("ANY ERROR DATABASE"))

	sut := productrepository.New(mock)
	products, err := sut.Fetch()

	require.NotNil(t, err)
	require.Nil(t, products)
}