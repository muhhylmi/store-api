package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
)

func (repository *UserRepositoryImpl) Save(ctx context.Context, user domain.Users) (domain.Users, error) {
	// SQL := "insert into Product(name) values (?)"
	// result, err := tx.ExecContext(ctx, SQL, Product.Name)
	// exception.PanicIfError(err)
	// id, err := result.LastInsertId()
	// exception.PanicIfError(err)
	// Product.Id = int(id)

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
	result := repository.DB.Gorm.Where(&domain.Users{Username: username}).First(&data)
	return data, result.Error
}