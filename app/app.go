package app

import "github.com/goblinlordx/go-test/accounting"

type App struct {
	AccountRepository     *accounting.AccountRepository
	CategoryRepository    *accounting.CategoryRepository
	TransactionRepository *accounting.TransactionRepository

	AccountCommander     *accounting.AccountCommander
	TransactionCommander *accounting.TransactionCommander
}

func ProvideApp(ar *accounting.AccountRepository, cr *accounting.CategoryRepository, tr *accounting.TransactionRepository, ac *accounting.AccountCommander, tc *accounting.TransactionCommander) App {
	return App{
		AccountRepository:     ar,
		CategoryRepository:    cr,
		TransactionRepository: tr,

		AccountCommander:     ac,
		TransactionCommander: tc,
	}
}
