package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
)

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product domain.Product) (domain.Product, error) {

	result := repository.DB.Gorm.Create(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	result := repository.DB.Gorm.Save(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	result := repository.DB.Gorm.Delete(&domain.Product{BaseModel: domain.BaseModel{
		ID: product.BaseModel.ID,
	}})
	return result.Error
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (*domain.Product, error) {
	var product *domain.Product
	result := repository.DB.Gorm.Where(&domain.Product{BaseModel: domain.BaseModel{
		ID: productId,
	}}).First(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []*domain.Product {
	var products []*domain.Product
	repository.DB.Gorm.Model(&domain.Product{}).Find(&products)
	return products
}
