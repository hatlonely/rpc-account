package mysql

import (
	"context"

	"github.com/pkg/errors"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-account/internal/storage"
)

func init() {
	storage.RegisterStorage("MySQL", NewMySQLWithOptions)
}

func NewMySQLWithOptions(options *wrap.GORMDBWrapperOptions, opts ...refx.Option) (*MySQL, error) {
	db, err := wrap.NewGORMDBWrapperWithOptions(options, opts...)

	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewGORMDBWrapperWithOptions failed")
	}

	if err := db.AutoMigrate(context.Background(), &storage.Account{}); err != nil {
		return nil, errors.WithMessage(err, "db.AutoMigrate failed")
	}

	return &MySQL{
		db: db,
	}, nil
}

type MySQL struct {
	db *wrap.GORMDBWrapper
}
