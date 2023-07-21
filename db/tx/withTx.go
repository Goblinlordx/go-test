package db_tx

import (
	"context"
	"database/sql"

	"github.com/goblinlordx/go-test/db"
	"gorm.io/gorm"
)

func WithTx(ctx context.Context, fn func(tx *gorm.DB) error, opts ...*sql.TxOptions) error {
	return db.WithDb(ctx, func(db *gorm.DB) error {
		return db.Transaction(fn, opts...)
	})
}
