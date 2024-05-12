package domain

// products ||--|{ categories: have_1
//     products{
//         string id
//         string product_name
//         string price
//         string category_id
//         timestamp created_at
//         bool is_deleted
//     }
// categories{
// 	string id
// 	string category_name
// 	timestamp created_at
// 	bool is_deleted
// }

type Categories struct {
	BaseModel

	CategoryName string `gorm:"column:category_name"`
}

func (Categories) TableName() string {
	return "categories"
}

type Product struct {
	BaseModel

	ProductName string     `gorm:"column:product_name"`
	Price       int64      `gorm:"column:price"`
	CategoryId  string     `gorm:"column:category_id"`
	Category    Categories `gorm:"foreignKey:CategoryId;references:ID"`
}

func (Product) TableName() string {
	return "products"
}
