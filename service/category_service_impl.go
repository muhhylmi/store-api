package service

import (
	"context"

	"go-store-api/model/domain"
	"go-store-api/model/web"
	"go-store-api/utils/exception"
	"go-store-api/utils/objects"
	"go-store-api/utils/wrapper"

	"github.com/google/uuid"
)

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	l := service.Logger.LogWithContext("product_service", "Create")

	err := service.Validate.Struct(request)
	if err != nil {
		l.Error(err)
		exception.PanicIfError(err)
	}
	category := domain.Categories{
		BaseModel: domain.BaseModel{
			ID:        uuid.NewString(),
			IsDeleted: objects.ToPointer(false),
			CreatedBy: &request.UserId,
		},
		CategoryName: request.CategoryName,
	}
	result, err := service.Repository.Save(ctx, category)
	if err != nil {
		l.Error(err)
		panic(wrapper.NewNotFoundError(err.Error()))
	}

	return web.ToCategoryRersponse(result)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	result := service.Repository.FindAll(ctx)
	return web.ToCategoryRersponses(result)
}
