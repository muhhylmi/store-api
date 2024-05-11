package service

import (
	"context"

	"github.com/muhhylmi/store-api/model/web"
)

type ProductService interface {
	Create(ctx context.Context, request web.ProductCreateRequest) web.ProductResponse
	Update(ctx context.Context, request web.ProductUpdateRequest) web.ProductResponse
	Delete(ctx context.Context, ProductId int)
	FindById(ctx context.Context, ProductId int) web.ProductResponse
	FindAll(ctx context.Context) []web.ProductResponse
}
