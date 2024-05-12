package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/utils/logger"
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
