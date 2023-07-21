package db

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type dbCtxKeyType string

const ctxKey = dbCtxKeyType("db")

func WithDbCtx(ctx context.Context, fn func(ctx context.Context) error) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	childCtx := context.WithValue(ctx, ctxKey, db)
	return fn(childCtx)
}

func WithDb(ctx context.Context, fn func(db *gorm.DB) error) error {
	c := ctx.Value(ctxKey)
	if c == nil {
		return errors.New("db not initialized")
	}
	v := c.(*gorm.DB)
	return fn(v)
}
