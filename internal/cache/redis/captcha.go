package redis

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/hatlonely/rpc-account/internal/cache"
	"github.com/pkg/errors"
)

func (r *Redis) GetOrSetCaptcha(ctx context.Context, key string) (string, error) {
	key = r.options.Prefix + key
	val, err := r.client.Get(ctx, key).Result()
	if err == nil {
		return val, nil
	}
	if err != redis.Nil {
		return "", errors.WithMessagef(err, "redis get key [%v] failed", key)
	}

	captcha := cache.GenerateCaptcha()
	if err := r.client.Set(ctx, key, captcha, r.options.CaptchaExpiration).Err(); err != nil {
		return "", errors.WithMessagef(err, "redis set key [%v] failed", key)
	}
	return captcha, nil
}

func (r *Redis) GetCaptcha(ctx context.Context, key string) (string, error) {
	key = r.options.Prefix + key
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", errors.WithMessagef(err, "redis get key [%v] failed", key)
	}
	return val, nil
}
