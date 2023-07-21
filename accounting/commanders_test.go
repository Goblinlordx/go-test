package accounting

import (
	"context"
	"testing"

	db_tx "github.com/goblinlordx/go-test/db/tx"
	"github.com/google/uuid"
)

func TestAccountCommander(t *testing.T) {
	err := db_tx.WithTestTx(func(ctx context.Context) error {
		var err error
		var id uuid.UUID

		ac := ProvideTestAccountCommander()

		createInput := AccountCreateInput{
			Name:           "Test 1",
			CurrencySymbol: "USD",
		}
		if id, err = ac.Create(ctx, createInput); err != nil {
			return err
		}

		nameUpdate := "Test 1 updated"
		currencyUpdate := "KRW"
		updateInput := AccountUpdateInput{
			Id:             id.String(),
			Name:           &nameUpdate,
			CurrencySymbol: &currencyUpdate,
		}
		if err = ac.Update(ctx, updateInput); err != nil {
			return err
		}

		deleteInput := AccountDeleteInput{
			Ids: []string{id.String()},
		}
		if err = ac.Delete(ctx, deleteInput); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestTransactionCommander(t *testing.T) {
	err := db_tx.WithTestTx(func(ctx context.Context) error {
		var err error
		var id uuid.UUID

		ac := ProvideTestTransactionCommander()

		createInput := TransactionCreateInput{
			Payee:  "Payee 1",
			Amount: 1,
			Note:   "Note",
		}
		if id, err = ac.Create(ctx, createInput); err != nil {
			return err
		}

		payeeUpdate := "Test 1 updated"
		amountUpdate := 2
		noteUpdate := "Updated Note"
		updateInput := TransactionUpdateInput{
			Id:     id.String(),
			Payee:  &payeeUpdate,
			Amount: &amountUpdate,
			Note:   &noteUpdate,
		}
		if err = ac.Update(ctx, updateInput); err != nil {
			return err
		}

		deleteInput := TransactionDeleteInput{
			Ids: []string{id.String()},
		}
		if err = ac.Delete(ctx, deleteInput); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		t.Fatal(err)
	}
}
