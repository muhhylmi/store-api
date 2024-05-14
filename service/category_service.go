package service

import (
	"context"

	"go-store-api/model/web"
	"go-store-api/repository"
	"go-store-api/utils/logger"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	Logger     *logger.Logger
	Repository repository.CategoryRepository
	Validate   *validator.Validate
}

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}

func NewCategoryService(logger *logger.Logger, category repository.CategoryRepository, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Logger:     logger,
		Repository: category,
		Validate:   validate,
	}
}
