package currency

import "github.com/google/wire"

var Module = wire.NewSet(ProvideCurrencyRepository)
