// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/goblinlordx/go-test/app"
)

// Injectors from wire.go:

func InitializeApp() app.App {
	appApp := app.ProvideApp()
	return appApp
}
