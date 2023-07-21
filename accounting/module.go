package accounting

import (
	"github.com/goblinlordx/go-test/currency"
	"github.com/google/wire"
)

var Module = wire.NewSet(ProvideAccountRepository, ProvideCategoryRepository, ProvideTransactionRepository, ProvideAccountCommander, ProvideTransactionCommander, currency.Module)
