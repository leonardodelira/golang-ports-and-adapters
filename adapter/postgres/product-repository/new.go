package productrepository

import (
	"go/cleanarch/adapter/postgres"
	"go/cleanarch/core/domain"
)

type repository struct {
    db postgres.PoolInterface
}

// New returns contract implementation of ProductRepository
func New(db postgres.PoolInterface) domain.ProductRepository {
    return &repository{
        db: db,
    }
}