package accounting

import (
	"context"
	"fmt"
	"testing"

	db_tx "github.com/goblinlordx/go-test/db/tx"
)

func TestAccountQuerier(t *testing.T) {
	err := db_tx.WithTestTx(func(ctx context.Context) error {
		filterer := ProvideTestAccountFilterer()
		repo := ProvideTestAccountRepository()

		ctx, err := filterer.Apply(ctx, AccountFilterInput{
			Id:             "9abad5e0-7f1f-430d-a319-407c40effb34",
			Name:           "scar",
			CurrencySymbol: "USD",
		})
		if err != nil {
			return err
		}

		res, err := repo.List(ctx)
		if err != nil {
			return err
		}

		if len(res) != 1 {
			return fmt.Errorf("expected results length: 1, got: %d", len(res))
		}

		// app.AccountRepository.Search(ctx)
		return nil
	})

	if err != nil {
		t.Error(err)
	}
}

func TestTransactionQuerier(t *testing.T) {

}
