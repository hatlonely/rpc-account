package redis

import (
	"time"

	"github.com/pkg/errors"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-account/internal/cache"
)

func init() {
	cache.RegisterCache("Redis", NewRedisWithOptions)
}

type Options struct {
	RedisClientWrapper wrap.RedisClientWrapperOptions
	Prefix             string
	CaptchaExpiration  time.Duration
}

func NewRedisWithOptions(options *Options, opts ...refx.Option) (*Redis, error) {
	client, err := wrap.NewRedisClientWrapperWithOptions(&options.RedisClientWrapper, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewRedisClientWrapperWithOptions failed")
	}

	return &Redis{
		client:  client,
		options: options,
	}, nil
}

type Redis struct {
	client  *wrap.RedisClientWrapper
	options *Options
}
