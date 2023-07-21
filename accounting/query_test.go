package accounting

import (
	"context"
	"errors"
	"testing"
	"time"

	db_tx "github.com/goblinlordx/go-test/db/tx"
)

func TestTransactionQuery(t *testing.T) {
	err := db_tx.WithTestTx(func(ctx context.Context) error {

		tr := TransactionRepository{}

		start := time.Date(
			2020, 01, 01, 0, 0, 0, 0, time.UTC)
		end := time.Date(
			2024, 01, 01, 0, 0, 0, 0, time.UTC)
		res, err := tr.AggregateByRange(ctx, RANGE_DAILY, start, end)
		if err != nil {
			return err
		}
		if len(res) == 0 {
			return errors.New("unexpected empty results set")
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
