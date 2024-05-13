package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
)

func (repository *ShoppingCartRepositoryImpl) Save(ctx context.Context, carts []domain.ShoppingCarts) ([]domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.CreateInBatches(&carts, len(carts))
	return carts, result.Error
}

func (repository *ShoppingCartRepositoryImpl) FindAll(ctx context.Context, req web.ListCartRequest) []*domain.ShoppingCarts {
	var carts []*domain.ShoppingCarts
	tx := repository.DB.Gorm.Model(&domain.ShoppingCarts{}).
		Preload("Product").
		Where("user_id = ? AND is_deleted = ?", req.AuthData.UserId, false)
	if req.Status != "" {
		tx.Where("status = ?", req.Status)
	}
	tx.Find(&carts)
	return carts
}

func (repository *ShoppingCartRepositoryImpl) FindById(ctx context.Context, Id string) (*domain.ShoppingCarts, error) {
	var cart *domain.ShoppingCarts
	result := repository.DB.Gorm.Where(&domain.ShoppingCarts{BaseModel: domain.BaseModel{
		ID: Id,
	}}).First(&cart)
	return cart, result.Error
}

func (repository *ShoppingCartRepositoryImpl) Update(ctx context.Context, cart domain.ShoppingCarts) (domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.Save(cart)
	return cart, result.Error
}
