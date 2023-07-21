//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func InitializeApp() App {
	wire.Build(Module)
	return App{}
}
