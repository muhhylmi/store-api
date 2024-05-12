package domain

// users{
// 	string id
// 	string username
// 	string password
// 	string role
// 	timestamp created_at
// 	bool is_deleted
// }
type Users struct {
	BaseModel

	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
}

func (Users) TableName() string {
	return "users"
}
