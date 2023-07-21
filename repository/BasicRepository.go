package repository

import (
	"context"
	"errors"

	"github.com/goblinlordx/go-test/db"
	db_tx "github.com/goblinlordx/go-test/db/tx"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BasicRepository[T any] struct{}

func (r BasicRepository[T]) GetById(ctx context.Context, dest *T, id uuid.UUID) error {
	err := r.FindById(ctx, dest, id)
	if dest == nil && err != nil {
		return errors.New("not found")
	}
	return err
}

func (r BasicRepository[T]) FindById(ctx context.Context, dest *T, id uuid.UUID) error {
	err := db.WithDb(ctx, func(db *gorm.DB) error {
		return db.First(dest, "id = ?", id.String()).Error
	})

	return err
}

func (r BasicRepository[T]) Save(ctx context.Context, ar []T) error {
	return db_tx.WithTx(ctx, func(tx *gorm.DB) error {
		for _, a := range ar {
			err := tx.Save(a).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (r BasicRepository[T]) Delete(ctx context.Context, records ...T) error {
	return db_tx.WithTx(ctx, func(tx *gorm.DB) error {
		for _, record := range records {
			err := tx.Delete(&record).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
