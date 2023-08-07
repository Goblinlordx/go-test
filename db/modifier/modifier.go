package modifier

import (
	"context"

	"github.com/goblinlordx/go-test/error_utils"
	"gorm.io/gorm"
)

type Modifier func(*gorm.DB) (*gorm.DB, error)

func Apply(db *gorm.DB, mr ...Modifier) error {
	c := error_utils.NewErrorChain()
	_ = c
	for _, m := range mr {
		c.Next(func() error {
			var err error
			db, err = m(db)
			return err
		})
	}
	return c.Check()
}

type modifierContextKey string

var mKey = modifierContextKey("modifiers")

func WithModifierScope(ctx context.Context, fn func(context.Context) error) error {
	ctx = context.WithValue(ctx, mKey, &[]Modifier{})
	return fn(ctx)
}

func AddModifier(ctx context.Context, m Modifier) context.Context {
	ms := GetModifiers(ctx)

	ms = append(ms, m)
	return context.WithValue(ctx, mKey, &ms)
}

func GetModifiers(ctx context.Context) []Modifier {
	m, ok := ctx.Value(mKey).(*[]Modifier)
	if !ok {
		return []Modifier{}
	}

	return *m
}
