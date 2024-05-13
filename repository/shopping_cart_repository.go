package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/databases"
	"github.com/muhhylmi/store-api/utils/logger"
)

type ShoppingCartRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type ShoppingCartRepository interface {
	Save(ctx context.Context, product domain.ShoppingCarts) (domain.ShoppingCarts, error)
	FindAll(ctx context.Context) []*domain.ShoppingCarts
	FindById(ctx context.Context, Id string) (*domain.ShoppingCarts, error)
}

func NewShoppingCartRepository(logger *logger.Logger, db *databases.DBService) ShoppingCartRepository {
	return &ShoppingCartRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
