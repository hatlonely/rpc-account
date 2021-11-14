package redis

import (
	"github.com/pkg/errors"

	"github.com/hatlonely/go-kit/refx"
	"github.com/hatlonely/go-kit/wrap"
	"github.com/hatlonely/rpc-account/internal/cache"
)

func init() {
	cache.RegisterStorage("MySQL", NewMySQLWithOptions)
}

func NewMySQLWithOptions(options *wrap.RedisClientWrapperOptions, opts ...refx.Option) (*MySQL, error) {
	client, err := wrap.NewRedisClientWrapperWithOptions(options, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "wrap.NewRedisClientWrapperWithOptions failed")
	}

	return &MySQL{
		client: client,
	}, nil
}

type MySQL struct {
	client *wrap.RedisClientWrapper
}
