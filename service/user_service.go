package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/utils/config"
	"github.com/muhhylmi/store-api/utils/logger"
)

type UserServiceImpl struct {
	Logger     *logger.Logger
	Repository repository.UserRepository
	Validate   *validator.Validate
	Config     *config.Configurations
}

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	// Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	// Delete(ctx context.Context, ProductId string)
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	// FindAll(ctx context.Context) []web.ProductResponse
}

func NewUserService(logger *logger.Logger, config *config.Configurations, repository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Logger:     logger,
		Repository: repository,
		Validate:   validate,
		Config:     config,
	}
}
