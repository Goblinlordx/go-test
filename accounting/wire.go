//go:build wireinject
// +build wireinject

package accounting

import "github.com/google/wire"

func ProvideTestAccountCommander() *AccountCommander {
	wire.Build(Module)

	return &AccountCommander{}
}

func ProvideTestTransactionCommander() *TransactionCommander {
	wire.Build(Module)

	return &TransactionCommander{}
}
