package migrate

import (
	"context"

	"github.com/goblinlordx/go-test/accounting"
	db_tx "github.com/goblinlordx/go-test/db/tx"
	"gorm.io/gorm"
)

func Up(ctx context.Context) error {
	return db_tx.WithTx(ctx, func(tx *gorm.DB) error {
		tx.AutoMigrate(&accounting.Account{}, &accounting.Transaction{}, &accounting.Category{})
		return nil
	})

}
