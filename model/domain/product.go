package domain

type Product struct {
	Id   string `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name"`
}
