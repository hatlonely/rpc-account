package mysql

import (
	"context"

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
		return nil, err
	}

	if !db.HasTable(&storage.Account{}) {
		if err := db.
			Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
			CreateTable(context.Background(), &storage.Account{}).
			Unwrap().Error; err != nil {
			return nil, err
		}
	} else {
		if err := db.AutoMigrate(context.Background(), &storage.Account{}).Unwrap().Error; err != nil {
			return nil, err
		}
	}

	return &MySQL{
		db: db,
	}, nil
}

type MySQL struct {
	db *wrap.GORMDBWrapper
}
