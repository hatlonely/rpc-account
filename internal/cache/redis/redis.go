package redis

import (
	"github.com/pkg/errors"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-account/internal/cache"
)

func init() {
	cache.RegisterCache("Redis", NewRedisWithOptions)
}

func NewRedisWithOptions(options *wrap.RedisClientWrapperOptions, opts ...refx.Option) (*Redis, error) {
	client, err := wrap.NewRedisClientWrapperWithOptions(options, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewRedisClientWrapperWithOptions failed")
	}

	return &Redis{
		client: client,
	}, nil
}

type Redis struct {
	client *wrap.RedisClientWrapper
}
