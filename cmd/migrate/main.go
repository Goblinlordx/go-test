package main

import (
	"context"
	"fmt"

	"github.com/goblinlordx/go-test/db"
	"github.com/goblinlordx/go-test/db/migrate"
)

func main() {
	fmt.Println("Starting migrations...")
	ctx := context.Background()
	err := db.WithDbCtx(ctx, func(ctx context.Context) error {
		return migrate.Up(ctx)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Migrations successful")
}
