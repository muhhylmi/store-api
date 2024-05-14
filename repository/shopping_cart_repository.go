package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/databases"
	"github.com/muhhylmi/store-api/utils/logger"
)

type ShoppingCartRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type ShoppingCartRepository interface {
	Save(ctx context.Context, carts []domain.ShoppingCarts) ([]domain.ShoppingCarts, error)
	FindAll(ctx context.Context, req web.ListCartRequest) []*domain.ShoppingCarts
	FindById(ctx context.Context, Id string) (*domain.ShoppingCarts, error)
	Update(ctx context.Context, cart domain.ShoppingCarts) (domain.ShoppingCarts, error)
	FindPendingByIds(ctx context.Context, Ids []string, auth web.AuthData) ([]domain.ShoppingCarts, error)
	UpdateByIds(ctx context.Context, Ids []string, cart domain.ShoppingCarts) (domain.ShoppingCarts, error)
}

func NewShoppingCartRepository(logger *logger.Logger, db *databases.DBService) ShoppingCartRepository {
	return &ShoppingCartRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
