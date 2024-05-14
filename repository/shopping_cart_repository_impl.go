package repository

import (
	"context"
	"errors"

	"go-store-api/model/domain"
	"go-store-api/model/web"
	"go-store-api/utils/objects"
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
		ID:        Id,
		IsDeleted: objects.ToPointer(false),
	}}).First(&cart)
	return cart, result.Error
}

func (repository *ShoppingCartRepositoryImpl) Update(ctx context.Context, cart domain.ShoppingCarts) (domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.Save(cart)
	return cart, result.Error
}

func (repository *ShoppingCartRepositoryImpl) FindPendingByIds(ctx context.Context, Ids []string, auth web.AuthData) ([]domain.ShoppingCarts, error) {
	var cart []domain.ShoppingCarts
	result := repository.DB.Gorm.Where(&domain.ShoppingCarts{}).
		Preload("Product").
		Where("is_deleted = ? AND id IN (?) AND user_id = ? AND status = ?", false, Ids, auth.UserId, web.PENDING_CART).
		Find(&cart)
	return cart, result.Error
}

func (repository *ShoppingCartRepositoryImpl) UpdateByIds(ctx context.Context, Ids []string, cart domain.ShoppingCarts) (domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.Model(&domain.ShoppingCarts{}).
		Where("id IN (?)", Ids).
		Updates(&domain.ShoppingCarts{
			BaseModel: domain.BaseModel{
				UpdatedBy: cart.UpdatedBy,
				UpdatedAt: cart.UpdatedAt,
			},
			Status: cart.Status,
		})
	if result.RowsAffected == 0 {
		return cart, errors.New("no update found")
	}
	return cart, result.Error
}
