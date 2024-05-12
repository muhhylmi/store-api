package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/databases"
	"github.com/muhhylmi/store-api/utils/logger"
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
