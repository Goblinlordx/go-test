//go:build wireinject
// +build wireinject

package accounting

import "github.com/google/wire"

func ProvideTestAccountCommander() *AccountCommander {
	wire.Build(Module)

	return &AccountCommander{}
}

func ProvideTestAccountRepository() *AccountRepository {
	wire.Build(Module)

	return &AccountRepository{}
}

func ProvideTestAccountFilterer() *AccountFilterer {
	wire.Build(Module)

	return &AccountFilterer{}
}

func ProvideTestTransactionCommander() *TransactionCommander {
	wire.Build(Module)

	return &TransactionCommander{}
}
