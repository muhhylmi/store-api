package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/objects"
)

func (repository *ShoppingCartRepositoryImpl) Save(ctx context.Context, cart domain.ShoppingCarts) (domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.Create(&cart)
	return cart, result.Error
}

func (repository *ShoppingCartRepositoryImpl) FindAll(ctx context.Context) []*domain.ShoppingCarts {
	var carts []*domain.ShoppingCarts
	repository.DB.Gorm.Where(&domain.ShoppingCarts{BaseModel: domain.BaseModel{
		IsDeleted: objects.ToPointer(false),
	}}).Find(&carts)
	return carts
}

func (repository *ShoppingCartRepositoryImpl) FindById(ctx context.Context, Id string) (*domain.ShoppingCarts, error) {
	var cart *domain.ShoppingCarts
	result := repository.DB.Gorm.Where(&domain.ShoppingCarts{BaseModel: domain.BaseModel{
		ID: Id,
	}}).First(&cart)
	return cart, result.Error
}
