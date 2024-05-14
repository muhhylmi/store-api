package repository

import (
	"context"

	"go-store-api/model/domain"
	"go-store-api/model/web"
	"go-store-api/utils/databases"
	"go-store-api/utils/logger"
)

type ProductRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type ProductRepository interface {
	Save(ctx context.Context, Product domain.Product) (domain.Product, error)
	Update(ctx context.Context, Product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, Product domain.Product) error
	FindById(ctx context.Context, ProductId string) (*web.ProductResponse, error)
	FindProductById(ctx context.Context, Id string) (*domain.Product, error)
	FindAll(ctx context.Context, req web.ProductListRequest) []*web.ProductResponse
	CountByIds(ctx context.Context, req []string) int64
}

func NewProductRepository(logger *logger.Logger, db *databases.DBService) ProductRepository {
	return &ProductRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
