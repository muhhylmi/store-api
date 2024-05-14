package service

import (
	"context"
	"time"

	"go-store-api/model/domain"
	"go-store-api/model/web"
	"go-store-api/utils/exception"
	"go-store-api/utils/objects"
	"go-store-api/utils/wrapper"

	"github.com/google/uuid"
)

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	_, errCheck := service.CategoryRepository.FindById(ctx, request.CategoryId)
	if errCheck != nil {
		l.Error(errCheck)
		panic(wrapper.NewNotFoundError("category not found"))
	}

	product := domain.Product{
		BaseModel: domain.BaseModel{
			ID:        uuid.NewString(),
			IsDeleted: objects.ToPointer(false),
			CreatedBy: &request.UserId,
		},
		Price:       request.Price,
		CategoryId:  request.CategoryId,
		ProductName: request.Name,
	}
	result, err := service.ProductRepository.Save(ctx, product)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToProductRersponse(result)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "Update")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	product, err := service.ProductRepository.FindProductById(ctx, request.Id)
	if err != nil {
		l.Error("product not found")
		panic(wrapper.NewNotFoundError("product not found"))
	}

	_, errCheckCat := service.CategoryRepository.FindById(ctx, request.CategoryId)
	if errCheckCat != nil {
		l.Error(errCheckCat)
		panic(wrapper.NewNotFoundError("category not found"))
	}

	product.ProductName = request.Name
	product.Price = request.Price
	product.CategoryId = request.CategoryId
	product.UpdatedAt = time.Now().Unix()
	product.UpdatedBy = request.AuthData.UserId

	return web.ToProductRersponse(*product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, request web.DeleteProductRequest) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "Delete")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}

	product, err := service.ProductRepository.FindProductById(ctx, request.Id)
	if err != nil {
		l.Error("product not found")
		panic(wrapper.NewNotFoundError("product not found"))
	}

	product.IsDeleted = objects.ToPointer(true)
	product.UpdatedAt = time.Now().Unix()
	product.UpdatedBy = request.AuthData.UserId

	service.ProductRepository.Update(ctx, *product)

	return web.ToProductRersponse(*product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "FindById")
	result, err := service.ProductRepository.FindById(ctx, productId)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return *result
}

func (service *ProductServiceImpl) FindAll(ctx context.Context, request web.ProductListRequest) []*web.ProductResponse {
	l := service.Logger.LogWithContext("service", "findAll")
	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	categories := service.ProductRepository.FindAll(ctx, request)

	return categories
}
