package redis

import (
	"context"

	"github.com/go-redis/redis"

	"github.com/hatlonely/rpc-account/internal/cache"
	"github.com/hatlonely/rpc-account/internal/storage"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

func (r *Redis) SetToken(ctx context.Context, account *storage.Account) (string, error) {
	token := cache.GenerateToken()
	key := r.options.Prefix + token
	val, _ := jsoniter.MarshalToString(account)
	if err := r.client.Set(ctx, key, val, r.options.TokenExpiration).Err(); err != nil {
		return "", errors.Wrapf(err, "client.Get failed. key: [%s]", key)
	}
	return token, nil
}

func (r *Redis) GetToken(ctx context.Context, token string) (*storage.Account, error) {
	key := r.options.Prefix + token
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, cache.ErrNotFound
	}
	if err != nil {
		return nil, errors.Wrapf(err, "client.Get failed. key: [%s]", key)
	}
	var account storage.Account
	if err := jsoniter.UnmarshalFromString(val, &account); err != nil {
		return nil, errors.Wrapf(err, "jsoniter.UnmarshalFromString failed. key: [%s]", key)
	}
	return &account, nil
}

func (r *Redis) DelToken(ctx context.Context, token string) error {
	key := r.options.Prefix + token
	if err := r.client.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "client.Del failed")
	}

	return nil
}
