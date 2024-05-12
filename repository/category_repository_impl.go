package repository

import (
	"context"
	"fmt"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/utils/objects"
)

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, category domain.Categories) (domain.Categories, error) {
	result := repository.DB.Gorm.Create(&category)
	return category, result.Error
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context) []*domain.Categories {
	var caategories []*domain.Categories
	repository.DB.Gorm.Where(&domain.Product{BaseModel: domain.BaseModel{
		IsDeleted: objects.ToPointer(false),
	}}).Find(&caategories)
	return caategories
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, Id string) (*domain.Categories, error) {
	fmt.Println("============= Category Repo")

	var category *domain.Categories
	result := repository.DB.Gorm.Where(&domain.Categories{BaseModel: domain.BaseModel{
		ID: Id,
	}}).First(&category)
	return category, result.Error
}
