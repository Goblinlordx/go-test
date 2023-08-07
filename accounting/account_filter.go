package accounting

import (
	"context"
	"strings"

	"github.com/goblinlordx/go-test/currency"
	"github.com/goblinlordx/go-test/db/modifier"
	"github.com/google/uuid"
)

type AccountFilterInput struct {
	Id             string
	Name           string
	CurrencySymbol string
}

type AccountFilterer struct {
	cr *currency.CurrencyRepository
}

func ProvideAccountFilterer(cr *currency.CurrencyRepository) *AccountFilterer {
	return &AccountFilterer{
		cr: cr,
	}
}

func (af *AccountFilterer) Apply(ctx context.Context, input AccountFilterInput) (context.Context, error) {
	if input.Id != "" {
		id, err := uuid.Parse(input.Id)
		if err != nil {
			return ctx, err
		}
		ctx = modifier.StrEqual(ctx, "id", id.String())
	}
	if input.Name != "" {
		ctx = modifier.StrMatch(ctx, "name", input.Name)
	}
	if input.CurrencySymbol != "" {
		c, err := af.cr.GetBySymbol(strings.ToUpper(input.CurrencySymbol))
		if err != nil {
			return ctx, err
		}
		ctx = modifier.StrEqual(ctx, "currency_symbol", string(c.Symbol))
	}

	return ctx, nil
}
