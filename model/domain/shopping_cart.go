package domain

// shopping_carts ||--|{ shopping_cart_items: contains
//     users||--o{ shopping_carts: have
//     shopping_carts {
//         string id
//         string user_id
//         string status
//         timestamp created_at
//         bool is_deleted
//     }
//     shopping_cart_items ||--|{ products: have_1
//     shopping_cart_items{
//         string id
//         string product_id
//         int quantity
//     }
type ShoppingCarts struct {
	BaseModel

	UserId    string  `gorm:"column:user_id"`
	Status    string  `gorm:"column:status"`
	ProductId string  `gorm:"column:product_id"`
	Qty       int     `gorm:"column:qty"`
	Product   Product `gorm:"foreignKey:ProductId;references:ID"`
	User      Users   `json:"user" gorm:"foreignKey:UserId;references:ID"`
}

func (ShoppingCarts) TableName() string {
	return "shopping_carts"
}
