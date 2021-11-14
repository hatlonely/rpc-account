package storage

import (
	"context"
	"reflect"
	"time"

	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"
)

type Account struct {
	ID       string    `gorm:"type:char(32);primary_key" json:"id"`
	Email    string    `gorm:"type:varchar(64);unique_index:email_idx" json:"email"`
	Phone    string    `gorm:"type:varchar(64);unique_index:phone_idx" json:"phone"`
	Name     string    `gorm:"type:varchar(32);unique_index:name_idx" json:"name"`
	Password string    `gorm:"type:varchar(32)" json:"password"`
	Birthday time.Time `gorm:"type:timestamp default '1970-01-02 00:00:00'" json:"birthday"`
	Gender   int       `gorm:"type:int(1)" json:"gender"`
	Avatar   string    `gorm:"type:varchar(512)" json:"avatar"`
}

type Storage interface {
	PutAccount(ctx context.Context, article *Account) (string, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	UpdateAccount(ctx context.Context, article *Account) error
	DelAccount(ctx context.Context, id string) error
}

func RegisterStorage(key string, constructor interface{}) {
	refx.Register(reflect.TypeOf((*Storage)(nil)).Elem(), key, constructor)
}

func NewStorageWithOptions(options *refx.TypeOptions, opts ...refx.Option) (Storage, error) {
	v, err := refx.New(reflect.TypeOf((*Storage)(nil)).Elem(), options, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "refx.New failed")
	}
	return v.(Storage), nil
}
