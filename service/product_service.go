package service

import (
	"context"

	"go-store-api/model/web"
	"go-store-api/repository"
	"go-store-api/utils/logger"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	Logger             *logger.Logger
	ProductRepository  repository.ProductRepository
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, req web.DeleteProductRequest) web.ProductResponse
	FindById(ctx context.Context, ProductId string) web.ProductResponse
	FindAll(ctx context.Context, req web.ProductListRequest) []*web.ProductResponse
}

func NewProductService(logger *logger.Logger, ProductRepository repository.ProductRepository,
	categoryRepo repository.CategoryRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		Logger:             logger,
		ProductRepository:  ProductRepository,
		CategoryRepository: categoryRepo,
		Validate:           validate,
	}
}
