package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/objects"
	"github.com/muhhylmi/store-api/utils/wrapper"
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
