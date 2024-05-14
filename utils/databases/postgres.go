package databases

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var PostgresDB *gorm.DB

func InitPostgres(params *DBServiceVar) (*gorm.DB, error) {
	log := params.Logger.LogWithContext("dbConnection", "InitPostgres")
	dsn := params.PostgresUri
	PostgresDB, err := gorm.Open(postgres.Open(*dsn), &gorm.Config{
		Logger: logger.New(
			params.Logger.Logger,
			logger.Config{
				SlowThreshold:             100 * time.Millisecond,
				LogLevel:                  logger.Info,
				Colorful:                  true,
				IgnoreRecordNotFoundError: false,
				ParameterizedQueries:      false,
			},
		),
	})

	if err != nil {
		log.Info("Connection Postgres is Failed")
		return nil, err
	}
	log.Info("Success connect to postgres database")
	return PostgresDB, nil
}
