package service

import (
	"html/template"
	"time"

	"github.com/go-redis/redis"
	"github.com/hatlonely/go-kit/cli"
	"github.com/hatlonely/go-kit/kv"
	"github.com/jinzhu/gorm"

	"github.com/hatlonely/go-rpc/rpc-account/internal/storage"
)

type AccountService struct {
	mysqlCli *gorm.DB
	redisCli *redis.Client
	emailCli *cli.EmailCli
	kv       *kv.KV

	options         *Options
	captchaEmailTpl *template.Template
}

func NewAccountService(mysqlCli *gorm.DB, redisCli *redis.Client, emailCli *cli.EmailCli, opts ...Option) (*AccountService, error) {
	options := defaultOptions
	for _, opt := range opts {
		opt(&options)
	}

	return NewAccountServiceWithOptions(mysqlCli, redisCli, emailCli, &options)
}

func NewAccountServiceWithOptions(mysqlCli *gorm.DB, redisCli *redis.Client, emailCli *cli.EmailCli, options *Options) (*AccountService, error) {
	if !mysqlCli.HasTable(&storage.Account{}) {
		if err := mysqlCli.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&storage.Account{}).Error; err != nil {
			return nil, err
		}
	}

	captchaEmailTpl, err := template.New("captcha").Parse(captchaTpl)
	if err != nil {
		return nil, err
	}

	store, err := kv.NewRedisStore(redisCli, kv.WithRedisExpiration(options.AccountExpiration))
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

type Options struct {
	CaptchaExpiration time.Duration
	AccountExpiration time.Duration
}

var defaultOptions = Options{
	CaptchaExpiration: 5 * time.Minute,
	AccountExpiration: 30 * time.Minute,
}

type Option func(options *Options)

func WithCaptchaExpiration(captchaExpiration time.Duration) Option {
	return func(options *Options) {
		options.CaptchaExpiration = captchaExpiration
	}
}

func WithAccountExpiration(accountExpiration time.Duration) Option {
	return func(options *Options) {
		options.AccountExpiration = accountExpiration
	}
}
