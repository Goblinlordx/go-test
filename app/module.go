package app

import (
	"github.com/goblinlordx/go-test/accounting"
	"github.com/google/wire"
)

var Module = wire.NewSet(accounting.Module, ProvideApp)
