package domain

type BaseModel struct {
	ID string `gorm:"type:varchar(36);primaryKey" json:"id"` // uuid for identification

	IsDeleted *bool   `gorm:"type:boolean" json:"isDeleted"`                         // flag for soft delete
	CreatedAt int64   `gorm:"autoCreateTime;->;<-:create;" json:"createdAt"`         // unix time in seconds
	CreatedBy *string `gorm:"type:varchar(36);index;->;<-:create;" json:"createdBy"` // id of creator
	UpdatedAt int64   `gorm:"autoUpdateTime" json:"updatedAt"`                       // unix time in seconds
	UpdatedBy string  `gorm:"type:varchar(36);index" json:"updatedBy"`               // id of last updater
}
