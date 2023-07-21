package main

import (
	"context"

	"github.com/goblinlordx/go-test/accounting"
	"github.com/goblinlordx/go-test/app"
	"github.com/goblinlordx/go-test/db"
	"github.com/goblinlordx/go-test/error_utils"
)

func main() {
	app := app.InitializeApp()
	ctx := context.Background()
	err := db.WithDbCtx(ctx, func(ctx context.Context) error {
		atd := accounting.GetTestDataset()

		ec := error_utils.NewErrorChain()
		ec.Next(func() error { return app.AccountRepository.Save(ctx, atd.Accounts) })
		ec.Next(func() error { return app.CategoryRepository.Save(ctx, atd.Categories) })
		ec.Next(func() error { return app.TransactionRepository.Save(ctx, atd.Transactions) })

		return ec.Check()
	})
	if err != nil {
		panic(err)
	}
}
