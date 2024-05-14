package repository

import (
	"context"

	"go-store-api/model/domain"
	"go-store-api/model/web"
	"go-store-api/utils/databases"
	"go-store-api/utils/logger"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Logger *logger.Logger
	DB     *databases.DBService
}

type UserRepository interface {
	Save(ctx context.Context, user domain.Users) (domain.Users, error)
	Login(ctx context.Context, request web.LoginRequest) (*domain.Users, error)
	FindByUsername(ctx context.Context, username string) (*domain.Users, error)
	FindById(ctx context.Context, Id string) (*domain.Users, error)

	AdjustUpBalance(ctx context.Context, req domain.Users) (domain.Users, error)
	BeginTransaction(ctx context.Context) (context.Context, *gorm.DB)
	CommitTransaction(ctx context.Context) error
	RollbackTransaction(ctx context.Context) error
}

func NewUserRepository(logger *logger.Logger, db *databases.DBService) UserRepository {
	return &UserRepositoryImpl{
		Logger: logger,
		DB:     db,
	}
}
