package databases

import (
	"github.com/muhhylmi/store-api/utils/logger"
	"gorm.io/gorm"
)

type DBServiceVar struct {
	Logger      *logger.Logger
	PostgresUri *string
}

type DBService struct {
	Gorm *gorm.DB
}
