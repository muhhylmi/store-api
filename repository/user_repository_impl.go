package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/objects"
	"gorm.io/gorm"
)

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	result := repository.DB.Gorm.Create(&user)
	return user, result.Error
}

func (repository *UserRepositoryImpl) Login(ctx context.Context, user web.LoginRequest) (*domain.Users, error) {
	var data *domain.Users
	result := repository.DB.Gorm.Where(&domain.Users{Username: user.Username}).First(&data)
	return data, result.Error
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, username string) (*domain.Users, error) {
	var data *domain.Users
	result := repository.DB.Gorm.Where(&domain.Users{Username: username, BaseModel: domain.BaseModel{
		IsDeleted: objects.ToPointer(false),
	}}).First(&data)
	return data, result.Error
}

func (repository *UserRepositoryImpl) TopUpBalance(ctx context.Context, cart domain.Users) (domain.Users, error) {
	result := repository.DB.Gorm.Save(cart)
	return cart, result.Error
}

func (r *UserRepositoryImpl) BeginTransaction(ctx context.Context) (context.Context, *gorm.DB) {
	c, tx := r.DB.BeginTransaction(ctx)
	return c, tx
}

func (r *UserRepositoryImpl) CommitTransaction(ctx context.Context) error {
	return r.DB.CommitTransaction(ctx).Error
}

func (r *UserRepositoryImpl) RollbackTransaction(ctx context.Context) error {
	tx := r.DB.GetTransaction(ctx)
	return tx.Rollback().Error
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, Id string) (*domain.Users, error) {
	var data *domain.Users
	result := repository.DB.Gorm.Where(&domain.Users{BaseModel: domain.BaseModel{
		ID:        Id,
		IsDeleted: objects.ToPointer(false),
	}}).First(&data)
	return data, result.Error
}
