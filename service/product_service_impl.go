package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/logger"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

type ProductServiceImpl struct {
	Logger            *logger.Logger
	ProductRepository repository.ProductRepository
	Validate          *validator.Validate
}

func NewProductService(logger *logger.Logger, ProductRepository repository.ProductRepository, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		Logger:            logger,
		ProductRepository: ProductRepository,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	product := domain.Product{
		BaseModel: domain.BaseModel{
			ID: uuid.NewString(),
		},
		ProductName: request.Name,
	}
	result, err := service.ProductRepository.Save(ctx, product)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToProductRersponse(result)
}

// func (service *ProductServiceImpl) Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse {
// 	err := service.Validate.Struct(request)
// 	wrapper.PanicIfError(err)

// 	tx, err := service.DB.Begin()
// 	wrapper.PanicIfError(err)
// 	defer databases.CommitOrRollback(tx)

// 	Product, err := service.ProductRepository.FindById(ctx, tx, request.Id)
// 	if err != nil {
// 		panic(wrapper.NewNotFoundError(err.Error()))
// 	}

// 	Product.Name = request.Name
// 	Product = service.ProductRepository.Update(ctx, tx, Product)

// 	return web.ToProductRersponse(Product)
// }

// func (service *ProductServiceImpl) Delete(ctx context.Context, productId string) {
// 	tx, err := service.DB.Begin()
// 	wrapper.PanicIfError(err)
// 	defer databases.CommitOrRollback(tx)

// 	Product, err := service.ProductRepository.FindById(ctx, tx, ProductId)
// 	if err != nil {
// 		panic(wrapper.NewNotFoundError(err.Error()))
// 	}

// 	service.ProductRepository.Delete(ctx, tx, Product)
// }

func (service *ProductServiceImpl) FindById(ctx context.Context, productId string) web.ProductResponse {
	l := service.Logger.LogWithContext("product_service", "FindById")
	result, err := service.ProductRepository.FindById(ctx, productId)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToProductRersponse(*result)
}

// func (service *ProductServiceImpl) FindAll(ctx context.Context) []web.ProductResponse {
// 	l := service.Logger.LogWithContext("service", "findAll")
// 	tx, err := service.DB.Begin()
// 	if err != nil {
// 		l.Error(err)
// 		wrapper.PanicIfError(err)
// 	}
// 	defer databases.CommitOrRollback(tx)

// 	categories := service.ProductRepository.FindAll(ctx, tx)

// 	return web.ToProductRersponses(categories)
// }
