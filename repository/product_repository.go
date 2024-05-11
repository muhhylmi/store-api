package repository

import (
	"context"
	"database/sql"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/logger"
)

type ProductRepositoryImpl struct {
	Logger *logger.Logger
}

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Update(ctx context.Context, tx *sql.Tx, Product domain.Product) domain.Product
	Delete(ctx context.Context, tx *sql.Tx, Product domain.Product)
	FindById(ctx context.Context, tx *sql.Tx, ProductId int) (domain.Product, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}

func NewCategoryRepository(logger *logger.Logger) ProductRepository {
	return &ProductRepositoryImpl{
		Logger: logger,
	}
}
