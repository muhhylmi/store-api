package databases

import (
	"database/sql"
	"time"

	"github.com/muhhylmi/store-api/utils/exception"
	"github.com/muhhylmi/store-api/utils/logger"
)

func NewDB(logger *logger.Logger) *sql.DB {
	l := logger.LogWithContext("database", "postgres")

	db, err := sql.Open("postgres", "postgresql://user:password@localhost/golang_store_db")
	if err != nil {
		l.Error(err)
	}

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		exception.PanicIfError(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		exception.PanicIfError(errorCommit)
	}
}
