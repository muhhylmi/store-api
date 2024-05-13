package repository

import (
	"context"

	"github.com/muhhylmi/store-api/model/domain"
	"github.com/muhhylmi/store-api/model/web"
)

func (repository *ShoppingCartRepositoryImpl) Save(ctx context.Context, carts []domain.ShoppingCarts) ([]domain.ShoppingCarts, error) {
	result := repository.DB.Gorm.CreateInBatches(&carts, len(carts))
	return carts, result.Error
}

func (repository *ShoppingCartRepositoryImpl) FindAll(ctx context.Context, req web.ListCartRequest) []*domain.ShoppingCarts {
	var carts []*domain.ShoppingCarts

	// productTableName := domain.Product{}.TableName()
	// ShoppingCartTableName := domain.ShoppingCarts{}.TableName()
	// cartItemsTableName := domain.ShoppingCartItems{}.TableName()
	// querySelect := fmt.Sprintf(`%[1]s.id id, %[1]s.product_name name, %[1]s.category_id category_id,
	// %[1]s.price price ,%[2]s.category_name category_name`, productTableName, categoryTableName)

	tx := repository.DB.Gorm.Model(&domain.ShoppingCarts{}).
		Preload("Product").
		Where("user_id = ? AND is_deleted = ?", req.AuthData.UserId, false)
	if req.Status != "" {
		tx.Where("status = ?", req.Status)
	}
	tx.Find(&carts)

	// repository.DB.Gorm.Where(&domain.ShoppingCarts{BaseModel: domain.BaseModel{
	// 	IsDeleted: objects.ToPointer(false),
	// }}).Find(&carts)
	return carts
}

func (repository *ShoppingCartRepositoryImpl) FindById(ctx context.Context, Id string) (*domain.ShoppingCarts, error) {
	var cart *domain.ShoppingCarts
	result := repository.DB.Gorm.Where(&domain.ShoppingCarts{BaseModel: domain.BaseModel{
		ID: Id,
	}}).First(&cart)
	return cart, result.Error
}
