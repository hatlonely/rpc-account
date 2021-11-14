package service

import (
	"context"
	"html/template"

	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/cache"
	_ "github.com/hatlonely/rpc-account/internal/cache/redis"
	"github.com/hatlonely/rpc-account/internal/storage"
	_ "github.com/hatlonely/rpc-account/internal/storage/mysql"
	"github.com/pkg/errors"
)

type Options struct {
	Cache   refx.TypeOptions
	Storage refx.TypeOptions
	Email   cli.EmailOptions
}

type AccountService struct {
	api.UnsafeAccountServiceServer

	options *Options

	cache           cache.Cache
	storage         storage.Storage
	emailCli        *cli.EmailCli
	captchaEmailTpl *template.Template
}

func NewAccountServiceWithOptions(options *Options, opts ...refx.Option) (*AccountService, error) {
	captchaEmailTpl, err := template.New("captcha").Parse(captchaTpl)
	if err != nil {
		return nil, errors.Wrapf(err, "template.New failed")
	}

	cache, err := cache.NewCacheWithOptions(&options.Cache, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "cache.NewCacheWithOptions failed")
	}
	storage, err := storage.NewStorageWithOptions(&options.Storage, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "storage.NewStorageWithOptions failed")
	}

	return &AccountService{
		captchaEmailTpl: captchaEmailTpl,
		options:         options,
		cache:           cache,
		storage:         storage,
	}, nil
}

func (s *AccountService) Ping(context.Context, *api.Empty) (*api.Empty, error) {
	return &api.Empty{}, nil
}
