package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/databases"
	"github.com/muhhylmi/store-api/utils/logger"
)

type UserRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type UserRepository interface {
	Save(ctx context.Context, product domain.Users) (domain.Users, error)
	Login(ctx context.Context, request web.LoginRequest) (*domain.Users, error)
	FindByUsername(ctx context.Context, username string) (*domain.Users, error)
	// Update(ctx context.Context, Product domain.Product) (domain.Product, error)
	// Delete(ctx context.Context, Product domain.Product) error
	// FindById(ctx context.Context, ProductId string) (*domain.Product, error)
	// FindAll(ctx context.Context) []*domain.Product
}

func NewUserRepository(logger *logger.Logger, db *databases.DBService) UserRepository {
	return &UserRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}