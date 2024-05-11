package repository

import (
	"context"
	"database/sql"

	"github.com/muhhylmi/store-api/model/domain"
)

type ProductRepositoryImpl struct{}

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, Product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, ProductId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}

func NewCategoryRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}
