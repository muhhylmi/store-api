package repository

import (
	"context"
	"fmt"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/objects"
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

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (*web.ProductResponse, error) {
	var product *web.ProductResponse
	productTableName := domain.Product{}.TableName()
	categoryTableName := domain.Categories{}.TableName()
	querySelect := fmt.Sprintf(`%[1]s.id id, %[1]s.product_name name, %[1]s.category_id category_id,
	%[1]s.price price ,%[2]s.category_name category_name`, productTableName, categoryTableName)

	result := repository.DB.Gorm.Model(&domain.Product{}).
		Select(querySelect).
		Joins(fmt.Sprintf("INNER JOIN %[2]s on %[1]s.category_id = %[2]s.id", productTableName, categoryTableName)).
		Where(fmt.Sprintf("%s.id = ? AND is_deleted = ?", productTableName), productId, false).
		First(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context, req web.ProductListRequest) []*web.ProductResponse {
	var products []*web.ProductResponse
	productTableName := domain.Product{}.TableName()
	categoryTableName := domain.Categories{}.TableName()
	querySelect := fmt.Sprintf(`%[1]s.id id, %[1]s.product_name name, %[1]s.category_id category_id,
	%[1]s.price price ,%[2]s.category_name category_name`, productTableName, categoryTableName)

	tx := repository.DB.Gorm.Model(&domain.Product{}).
		Select(querySelect).
		Where(fmt.Sprintf("%s.is_deleted = ?", productTableName), false).
		Joins(fmt.Sprintf("INNER JOIN %[2]s on %[1]s.category_id = %[2]s.id", productTableName, categoryTableName))

	if req.CategoryId != "" {
		tx.Where(fmt.Sprintf("%s.category_id = ?", productTableName), req.CategoryId)
	}
	if req.Keyword != "" {
		tx.Where(fmt.Sprintf("%s.product_name ILIKE ?", productTableName), req.Keyword)
	}
	tx.Find(&products)
	return products
}

func (repository *ProductRepositoryImpl) CountByIds(ctx context.Context, productIds []string) int64 {
	var result int64

	repository.DB.Gorm.Model(domain.Product{}).
		Where("id IN ?", productIds).
		Count(&result)

	return result
}

func (repository *ProductRepositoryImpl) FindProductById(ctx context.Context, Id string) (*domain.Product, error) {
	var product *domain.Product
	result := repository.DB.Gorm.Where(&domain.Product{BaseModel: domain.BaseModel{
		ID:        Id,
		IsDeleted: objects.ToPointer(false),
	}}).First(&product)
	return product, result.Error
}
