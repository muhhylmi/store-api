package service

import (
	"context"

	"go-store-api/model/web"
	"go-store-api/repository"
	"go-store-api/utils/config"
	"go-store-api/utils/logger"

	"github.com/go-playground/validator/v10"
)

type UserServiceImpl struct {
	Logger     *logger.Logger
	Repository repository.UserRepository
	Validate   *validator.Validate
	Config     *config.Configurations
}

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserResponse
	Login(ctx context.Context, request web.LoginRequest) web.LoginResponse
	TopUpBalance(ctx context.Context, request web.TopUpRequest) web.TopUpResponse
}

func NewUserService(logger *logger.Logger, config *config.Configurations, repository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Logger:     logger,
		Repository: repository,
		Validate:   validate,
		Config:     config,
	}
}
