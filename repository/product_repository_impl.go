package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
)

func (repository *ProductRepositoryImpl) Save(ctx context.Context, product domain.Product) (domain.Product, error) {
	// SQL := "insert into Product(name) values (?)"
	// result, err := tx.ExecContext(ctx, SQL, Product.Name)
	// exception.PanicIfError(err)
	// id, err := result.LastInsertId()
	// exception.PanicIfError(err)
	// Product.Id = int(id)

	result := repository.DB.Gorm.Create(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	// SQL := "update Product set name = ? where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, Product.Name, Product.Id)
	// exception.PanicIfError(err)

	result := repository.DB.Gorm.Save(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) Delete(ctx context.Context, product domain.Product) error {
	// SQL := "delete from Product where id = ?"
	// _, err := tx.ExecContext(ctx, SQL, Product.Id)
	// exception.PanicIfError(err)
	result := repository.DB.Gorm.Delete(&domain.Product{BaseModel: domain.BaseModel{
		ID: product.BaseModel.ID,
	}})
	return result.Error
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, productId string) (*domain.Product, error) {
	// SQL := "select id, name from Product where id = ?"
	// rows, err := tx.QueryContext(ctx, SQL, ProductId)
	// exception.PanicIfError(err)
	// defer rows.Close()

	// Product := domain.Product{}
	// if rows.Next() {
	// 	err := rows.Scan(&Product.Id, &Product.Name)
	// 	exception.PanicIfError(err)
	// 	return Product, nil
	// } else {
	// 	return Product, errors.New("product is not found")
	// }
	var product *domain.Product
	result := repository.DB.Gorm.Where(&domain.Product{BaseModel: domain.BaseModel{
		ID: productId,
	}}).First(&product)
	return product, result.Error
}

func (repository *ProductRepositoryImpl) FindAll(ctx context.Context) []*domain.Product {
	// SQL := "select id, name from Product"
	// rows, err := tx.QueryContext(ctx, SQL)
	// exception.PanicIfError(err)
	// defer rows.Close()

	// var categories []domain.Product
	// for rows.Next() {
	// 	Product := domain.Product{}
	// 	rows.Scan(&Product.Id, &Product.Name)
	// 	categories = append(categories, Product)
	// }

	// return categories

	var products []*domain.Product
	repository.DB.Gorm.Model(&domain.Product{}).Find(&products)
	return products
}
