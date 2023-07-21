package accounting

import (
	"context"
	"errors"

	"github.com/goblinlordx/go-test/currency"
	"github.com/google/uuid"
)

type AccountCommander struct {
	ar *AccountRepository
	tr *TransactionRepository
	cr *currency.CurrencyRepository
}

func ProvideAccountCommander(ar *AccountRepository, tr *TransactionRepository, cr *currency.CurrencyRepository) *AccountCommander {
	return &AccountCommander{ar: ar, tr: tr, cr: cr}
}

type AccountCreateInput struct {
	Name           string
	CurrencySymbol string
}

func (ac AccountCommander) Create(ctx context.Context, input AccountCreateInput) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	_, err = ac.cr.GetBySymbol(input.CurrencySymbol)
	if err != nil {
		return uuid.UUID{}, errors.New("invalid currency symbol")
	}

	r := Account{
		Id:             id,
		Name:           input.Name,
		CurrencySymbol: currency.CurrencySymbol(input.CurrencySymbol),
	}

	err = ac.ar.Save(ctx, []Account{r})

	return id, err
}

type AccountUpdateInput struct {
	Id             string
	Name           *string
	CurrencySymbol *string
}

func (ac AccountCommander) Update(ctx context.Context, input AccountUpdateInput) error {
	id, err := uuid.Parse(input.Id)
	if err != nil {
		return errors.New("invalid id")
	}

	r := Account{}
	err = ac.ar.GetById(ctx, &r, id)
	if err != nil {
		return err
	}

	if input.CurrencySymbol != nil {
		c, err := ac.cr.GetBySymbol(*input.CurrencySymbol)
		if err != nil {
			return errors.New("invalid currency symbol")
		}

		r.CurrencySymbol = c.Symbol
	}

	if input.Name != nil {
		r.Name = *input.Name
	}

	return ac.ar.Save(ctx, []Account{r})
}

type AccountDeleteInput struct {
	Ids []string
}

func (ac AccountCommander) Delete(ctx context.Context, input AccountDeleteInput) error {
	records := make([]Account, len(input.Ids))
	for i, id := range input.Ids {
		uuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		records[i] = Account{Id: uuid}
	}

	return ac.ar.Delete(ctx, records...)
}

type TransactionCommander struct {
	tr *TransactionRepository
}

func ProvideTransactionCommander(tr *TransactionRepository) *TransactionCommander {
	return &TransactionCommander{tr: tr}
}

type TransactionCreateInput struct {
	Payee  string
	Amount int
	Note   string
}

func (tc TransactionCommander) Create(ctx context.Context, input TransactionCreateInput) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	r := Transaction{
		Id:     id,
		Payee:  input.Payee,
		Amount: input.Amount,
		Note:   input.Note,
	}

	err = tc.tr.Save(ctx, []Transaction{r})

	return id, err
}

type TransactionUpdateInput struct {
	Id     string
	Payee  *string
	Amount *int
	Note   *string
}

func (tc TransactionCommander) Update(ctx context.Context, input TransactionUpdateInput) error {
	id, err := uuid.Parse(input.Id)
	if err != nil {
		return errors.New("invalid id")
	}

	r := Transaction{}
	err = tc.tr.GetById(ctx, &r, id)
	if err != nil {
		return err
	}

	if input.Payee != nil {
		r.Payee = *input.Payee
	}

	if input.Amount != nil {
		r.Amount = *input.Amount
	}

	if input.Note != nil {
		r.Note = *input.Note
	}

	return tc.tr.Save(ctx, []Transaction{r})
}

type TransactionDeleteInput struct {
	Ids []string
}

func (tc TransactionCommander) Delete(ctx context.Context, input TransactionDeleteInput) error {
	records := make([]Transaction, len(input.Ids))
	for i, id := range input.Ids {
		uuid, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		records[i] = Transaction{Id: uuid}
	}

	return tc.tr.Delete(ctx, records...)
}
