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
	UserRepo    repository.UserRepository
	Validate    *validator.Validate
}

type ShoppingCartService interface {
	Create(ctx context.Context, request web.ShopingCartCreateRequest) []web.ShopingCartResponse
	FindAll(ctx context.Context, req web.ListCartRequest) []web.ListCartResponse
	Update(ctx context.Context, req web.UpdateCartRequest) web.ShopingCartResponse
	Delete(ctx context.Context, req web.DeleteCartRequest) web.ShopingCartResponse
	Checkout(ctx context.Context, req web.CheckoutCartRequest) web.CheckoutResponse
}

func NewShoppingCartService(logger *logger.Logger, cartRepo repository.ShoppingCartRepository, productRepo repository.ProductRepository,
	userRepo repository.UserRepository, validate *validator.Validate) ShoppingCartService {
	return &ShoppingCartServiceImpl{
		Logger:      logger,
		Repository:  cartRepo,
		ProductRepo: productRepo,
		UserRepo:    userRepo,
		Validate:    validate,
	}
}
