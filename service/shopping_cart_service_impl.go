package service

import (
	"context"
	"time"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/objects"
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
		l.Error("shopping cart not found")
		panic(wrapper.NewNotFoundError("shopping cart not found"))
	}

	_, errProduct := service.ProductRepo.FindById(ctx, req.ProductId)
	if errProduct != nil {
		l.Error("product is not found")
		panic(wrapper.NewNotFoundError("product is not found"))
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

func (service *ShoppingCartServiceImpl) Delete(ctx context.Context, req web.DeleteCartRequest) web.ShopingCartResponse {
	l := service.Logger.LogWithContext("shopping_cart_service", "Update")

	err := service.Validate.Struct(req)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	shoppingCart, errCheckCart := service.Repository.FindById(ctx, req.ShoppingCartId)
	if errCheckCart != nil {
		l.Error("shopping cart not found")
		panic(wrapper.NewNotFoundError("shopping cart not found"))
	}

	shoppingCart.BaseModel.UpdatedBy = req.AuthData.UserId
	shoppingCart.BaseModel.UpdatedAt = time.Now().Unix()
	shoppingCart.IsDeleted = objects.ToPointer(true)

	result, errUpdate := service.Repository.Update(ctx, *shoppingCart)
	if errUpdate != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	return web.ToUpdateShopingCartResponse(result)
}

func (service *ShoppingCartServiceImpl) Checkout(ctx context.Context, req web.CheckoutCartRequest) web.CheckoutResponse {
	l := service.Logger.LogWithContext("shopping_cart_service", "Update")

	err := service.Validate.Struct(req)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	// begin transaction
	ctx, _ = service.UserRepo.BeginTransaction(ctx)

	shoppingCarts, _ := service.Repository.FindPendingByIds(ctx, req.ShoppingCartIds, req.AuthData)
	if len(shoppingCarts) != len(req.ShoppingCartIds) {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error("some shopping cart not found")
		panic(wrapper.NewNotFoundError("some shopping cart not found"))
	}
	totalPrice := web.GetTotalPrice(shoppingCarts)

	cart := domain.ShoppingCarts{
		BaseModel: domain.BaseModel{
			UpdatedAt: time.Now().Unix(),
			UpdatedBy: req.AuthData.UserId,
		},
		Status: web.SUCCESS_CART,
	}

	user, errUser := service.UserRepo.FindById(ctx, req.AuthData.UserId)
	if errUser != nil {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error(errUser.Error())
		panic(wrapper.NewNotFoundError(errUser.Error()))
	}
	if user.Balance < float64(totalPrice) {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error("user balance is not enough")
		panic(wrapper.NewStatuConflictError("user balance is not enough, please top up by admin"))
	}

	_, errUpdate := service.Repository.UpdateByIds(ctx, req.ShoppingCartIds, cart)
	if errUpdate != nil {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error(err)
		exception.PanicIfError(err)
	}

	// adjust user balance
	user.BaseModel.UpdatedBy = req.AuthData.UserId
	user.BaseModel.UpdatedAt = time.Now().Unix()
	user.Balance -= float64(totalPrice)
	if _, err := service.UserRepo.AdjustUpBalance(ctx, *user); err != nil {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error("cannot adjust balance")
		panic(wrapper.NewStatuConflictError("cannot make payment"))
	}

	// commit transaction
	if errCommit := service.UserRepo.CommitTransaction(ctx); errCommit != nil {
		service.UserRepo.RollbackTransaction(ctx)
		l.Error(errCommit.Error())
		panic(wrapper.NewStatuConflictError(errCommit.Error()))
	}

	return web.CheckoutResponse{
		ShoppingCartIds: req.ShoppingCartIds,
		Message:         "Success Checkout Cart",
	}
}
