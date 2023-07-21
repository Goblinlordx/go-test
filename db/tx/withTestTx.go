package db_tx

import (
	"context"
	"errors"

	"github.com/goblinlordx/go-test/db"
	"gorm.io/gorm"
)

func WithTestTx(fn func(ctx context.Context) error) error {
	ctx := context.Background()
	rbErr := errors.New("rollback")
	err := db.WithDbCtx(ctx, func(ctx context.Context) error {
		return WithTx(ctx, func(db *gorm.DB) error {
			txInnerErr := fn(ctx)
			if txInnerErr != nil {
				return txInnerErr
			}

			return rbErr
		})
	})
	if err != rbErr {
		return err
	}
	return nil
}
