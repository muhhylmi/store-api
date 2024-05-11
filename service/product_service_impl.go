package service

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/helper"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(ProductRepository repository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product := domain.Product{
		Name: request.Name,
	}
	Product = service.ProductRepository.Save(ctx, tx, Product)

	return web.ToProductRersponse(Product)
}

func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	Product.Name = request.Name
	Product = service.ProductRepository.Update(ctx, tx, Product)

	return web.ToProductRersponse(Product)
}

func (service *ProductServiceImpl) Delete(ctx context.Context, ProductId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	service.ProductRepository.Delete(ctx, tx, Product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, ProductId int) web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}

	return web.ToProductRersponse(Product)
}

func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.ProductRepository.FindAll(ctx, tx)

	return web.ToProductRersponses(categories)
}
