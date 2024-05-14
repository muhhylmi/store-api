package databases

import (
	"context"

	"gorm.io/gorm"
)

const TransactionContextKey = "postgres:transaction"

func (db *DBService) BeginTransaction(ctx context.Context) (context.Context, *gorm.DB) {
	tx := ctx.Value(TransactionContextKey)
	if tx == nil {
		tx = db.Gorm.Begin()
		ctx = context.WithValue(ctx, TransactionContextKey, tx)
	}

	return ctx, tx.(*gorm.DB)
}

func (db *DBService) GetTransaction(ctx context.Context) *gorm.DB {
	tx := ctx.Value(TransactionContextKey)
	if tx == nil {
		tx = db.Gorm
	}

	return tx.(*gorm.DB).WithContext(ctx)
}

func (db *DBService) RollbackTransaction(ctx context.Context) *gorm.DB {
	tx := ctx.Value(TransactionContextKey).(*gorm.DB)
	return tx.Rollback()
}

func (db *DBService) CommitTransaction(ctx context.Context) *gorm.DB {
	tx := ctx.Value(TransactionContextKey).(*gorm.DB)
	return tx.Commit()
}
