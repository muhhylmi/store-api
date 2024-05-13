package service

import (
	"context"
	"time"

	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func (service *ShoppingCartServiceImpl) Create(ctx context.Context, request web.ShopingCartCreateRequest) []web.ShopingCartResponse {
	l := service.Logger.LogWithContext("shopping_cart_service", "Create")

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
	l := service.Logger.LogWithContext("shopping_cart_service", "FindAll")

	err := service.Validate.Struct(req)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	result := service.Repository.FindAll(ctx, req)
	return web.ToCartListResponse(result)
}

func (service *ShoppingCartServiceImpl) Update(ctx context.Context, req web.UpdateCartRequest) web.ShopingCartResponse {
	l := service.Logger.LogWithContext("shopping_cart_service", "Update")

	err := service.Validate.Struct(req)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	shoppingCart, errCheckCart := service.Repository.FindById(ctx, req.ShoppingCartId)
	if errCheckCart != nil {
		l.Error(errCheckCart)
		panic(wrapper.NewNotFoundError(errCheckCart.Error()))
	}

	_, errProduct := service.ProductRepo.FindById(ctx, req.ProductId)
	if errProduct != nil {
		l.Error(errProduct)
		panic(wrapper.NewNotFoundError(errProduct.Error()))
	}
	shoppingCart.BaseModel.UpdatedBy = req.AuthData.UserId
	shoppingCart.BaseModel.UpdatedAt = time.Now().Unix()
	shoppingCart.ProductId = req.ProductId
	shoppingCart.Qty = req.Qty

	result, errUpdate := service.Repository.Update(ctx, *shoppingCart)
	if errUpdate != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	return web.ToUpdateShopingCartResponse(result)
}
