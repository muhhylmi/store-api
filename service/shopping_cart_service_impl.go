package service

import (
	"context"

	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func (service *ShoppingCartServiceImpl) Create(ctx context.Context, request web.ShopingCartCreateRequest) web.ShopingCartResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	productIds := web.ToProductIds(request.Items)
	countId := service.ProductRepo.CountByIds(ctx, productIds)
	if len(productIds) != int(countId) {
		l.Error("some product is not found")
		panic(wrapper.NewNotFoundError("some product is not found"))
	}

	shoppingCart := web.ToShoppingCartModel(request)
	result, err := service.Repository.Save(ctx, shoppingCart)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToShoopingCartRersponse(result)
}

func (service *ShoppingCartServiceImpl) FindAll(ctx context.Context, req web.ListCartRequest) []web.ListCartResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(req)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	result := service.Repository.FindAll(ctx, req)
	return web.ToCartListResponse(result)
}
