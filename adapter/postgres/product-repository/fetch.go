package productrepository

import (
	"context"
	"go/cleanarch/core/domain"
)

func (repository repository) Fetch() (*domain.Pagination[[]domain.Product], error) {
    ctx := context.Background()
    products := []domain.Product{}
    total := int32(0)
    
    result, err := repository.db.Query(ctx, "SELECT * FROM product")
    if err != nil {
        return nil, err
    }

    for result.Next() {
        var product domain.Product
        result.Scan(&product.ID, &product.Name, &product.Price, &product.Description)
        products = append(products, product)
    }

    return &domain.Pagination[[]domain.Product]{
        Items: products,
        Total: total,
    }, nil
}