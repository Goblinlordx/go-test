package modifier

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

func StrEqual(ctx context.Context, k string, v string) context.Context {
	return AddModifier(ctx, func(db *gorm.DB) (*gorm.DB, error) {
		db = db.Where(fmt.Sprintf("%s = ?", k), v)
		return db, nil
	})
}

func StrMatch(ctx context.Context, k string, s string) context.Context {
	return AddModifier(ctx, func(db *gorm.DB) (*gorm.DB, error) {
		db = db.Where(fmt.Sprintf("UPPER(%s) LIKE \"%%\" || UPPER(?) || \"%%\"", k), s)
		return db, nil
	})
}

func StrPrefixMatch(ctx context.Context, k string, s string) context.Context {
	return AddModifier(ctx, func(db *gorm.DB) (*gorm.DB, error) {
		db = db.Where(fmt.Sprintf("%s LIKE ? || \"%%\"", k), s)
		return db, nil
	})
}
