package repository

import (
	"context"

	"go-store-api/model/domain"
	"go-store-api/utils/databases"
	"go-store-api/utils/logger"
)

type CategoryRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type CategoryRepository interface {
	Save(ctx context.Context, product domain.Categories) (domain.Categories, error)
	FindAll(ctx context.Context) []*domain.Categories
	FindById(ctx context.Context, Id string) (*domain.Categories, error)
}

func NewCategoryRepository(logger *logger.Logger, db *databases.DBService) CategoryRepository {
	return &CategoryRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
