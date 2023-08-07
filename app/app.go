package app

import (
	"github.com/goblinlordx/go-test/accounting"
	"github.com/goblinlordx/go-test/currency"
)

type App struct {
	CurrencyRepository    *currency.CurrencyRepository
	AccountRepository     *accounting.AccountRepository
	CategoryRepository    *accounting.CategoryRepository
	TransactionRepository *accounting.TransactionRepository

	AccountFilterer      *accounting.AccountFilterer
	AccountCommander     *accounting.AccountCommander
	TransactionCommander *accounting.TransactionCommander
}

func ProvideApp(curr *currency.CurrencyRepository, ar *accounting.AccountRepository, cr *accounting.CategoryRepository, tr *accounting.TransactionRepository, af *accounting.AccountFilterer, ac *accounting.AccountCommander, tc *accounting.TransactionCommander) App {
	return App{
		CurrencyRepository:    curr,
		AccountRepository:     ar,
		CategoryRepository:    cr,
		TransactionRepository: tr,

		AccountFilterer:      af,
		AccountCommander:     ac,
		TransactionCommander: tc,
	}
}
