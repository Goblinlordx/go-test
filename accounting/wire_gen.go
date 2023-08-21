// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package accounting

import (
	"github.com/goblinlordx/go-test/currency"
)

// Injectors from wire.go:

func ProvideTestAccountCommander() *AccountCommander {
	accountRepository := ProvideAccountRepository()
	transactionRepository := ProvideTransactionRepository()
	currencyRepository := currency.ProvideCurrencyRepository()
	accountCommander := ProvideAccountCommander(accountRepository, transactionRepository, currencyRepository)
	return accountCommander
}

func ProvideTestAccountRepository() *AccountRepository {
	accountRepository := ProvideAccountRepository()
	return accountRepository
}

func ProvideTestAccountFilterer() *AccountFilterer {
	currencyRepository := currency.ProvideCurrencyRepository()
	accountFilterer := ProvideAccountFilterer(currencyRepository)
	return accountFilterer
}

func ProvideTestTransactionCommander() *TransactionCommander {
	transactionRepository := ProvideTransactionRepository()
	transactionCommander := ProvideTransactionCommander(transactionRepository)
	return transactionCommander
}
