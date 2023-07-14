//go:build wireinject
// +build wireinject

package main

import (
	"github.com/goblinlordx/go-test/app"
	"github.com/google/wire"
)

func InitializeApp() app.App {
	wire.Build(app.ProvideApp)
	return app.App{}
}
