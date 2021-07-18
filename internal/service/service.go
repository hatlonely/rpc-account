package service

import (
	"context"
	"html/template"
	"time"

	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/kv"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-account/api/gen/go/api"

	"github.com/hatlonely/rpc-account/internal/storage"
)

type Options struct {
	CaptchaExpiration time.Duration `dft:"5m"`
	AccountExpiration time.Duration `dft:"30m"`
}

type AccountService struct {
	api.UnsafeAccountServiceServer

	mysqlCli *wrap.GORMDBWrapper
	redisCli *wrap.RedisClientWrapper
	emailCli *cli.EmailCli
	kv       *kv.KV

	options         *Options
	captchaEmailTpl *template.Template
}

func NewAccountServiceWithOptions(mysqlCli *wrap.GORMDBWrapper, redisCli *wrap.RedisClientWrapper, emailCli *cli.EmailCli, options *Options) (*AccountService, error) {
	if !mysqlCli.HasTable(&storage.Account{}) {
		if err := mysqlCli.Set(context.Background(), "gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(context.Background(), &storage.Account{}).Unwrap().Error; err != nil {
			return nil, err
		}
	}

	captchaEmailTpl, err := template.New("captcha").Parse(captchaTpl)
	if err != nil {
		return nil, err
	}

	store, err := kv.NewRedisStoreWithOptions(redisCli, &kv.RedisStoreOptions{
		Expiration: options.AccountExpiration,
	})
	if err != nil {
		return nil, err
	}
	kv := kv.NewKV(store, kv.WithStringKey())

	return &AccountService{
		mysqlCli:        mysqlCli,
		redisCli:        redisCli,
		emailCli:        emailCli,
		captchaEmailTpl: captchaEmailTpl,
		kv:              kv,
		options:         options,
	}, nil
}
