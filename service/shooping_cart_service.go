package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/utils/logger"
)

type ShoppingCartServiceImpl struct {
	Logger      *logger.Logger
	Repository  repository.ShoppingCartRepository
	ProductRepo repository.ProductRepository
	Validate    *validator.Validate
}

type ShoppingCartService interface {
	Create(ctx context.Context, request web.ShopingCartCreateRequest) []web.ShopingCartResponse
	FindAll(ctx context.Context, req web.ListCartRequest) []web.ListCartResponse
	Update(ctx context.Context, req web.UpdateCartRequest) web.ShopingCartResponse
}

func NewShoppingCartService(logger *logger.Logger, cartRepo repository.ShoppingCartRepository, productRepo repository.ProductRepository,
	validate *validator.Validate) ShoppingCartService {
	return &ShoppingCartServiceImpl{
		Logger:      logger,
		Repository:  cartRepo,
		ProductRepo: productRepo,
		Validate:    validate,
	}
}
