package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"lisichkinuriy/delivery/internal/adapters/ports"
)

var _ ports.IUnitOfWork = &UnitOfWork{}

type txKey struct{}

type UnitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) (*UnitOfWork, error) {
	if db == nil {
		return nil, errors.New("db can't be nil")
	}
	return &UnitOfWork{db: db}, nil
}

func (u *UnitOfWork) Begin(ctx context.Context) context.Context {
	tx := u.db.Begin()
	return context.WithValue(ctx, txKey{}, tx)
}

func (u *UnitOfWork) Commit(ctx context.Context) error {
	tx := GetTxFromContext(ctx)
	if tx != nil {
		return tx.Commit().Error
	}
	return nil
}

func (u *UnitOfWork) Rollback(ctx context.Context) error {
	tx := GetTxFromContext(ctx)
	if tx != nil {
		return tx.Rollback().Error
	}
	return nil
}

func GetTxFromContext(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return nil
}
