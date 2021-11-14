package cache

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"reflect"

	"github.com/hatlonely/rpc-account/internal/storage"

	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var ErrNotFound = errors.New("NotFound")

type Cache interface {
	GetOrSetCaptcha(ctx context.Context, key string) (string, error)
	GetCaptcha(ctx context.Context, key string) (string, error)
	GetToken(ctx context.Context, key string) (*storage.Account, error)
	SetToken(ctx context.Context, account *storage.Account) (string, error)
	DelToken(ctx context.Context, token string) error
}

func RegisterCache(key string, constructor interface{}) {
	refx.Register(reflect.TypeOf((*Cache)(nil)).Elem(), key, constructor)
}

func NewCacheWithOptions(options *refx.TypeOptions, opts ...refx.Option) (Cache, error) {
	v, err := refx.New(reflect.TypeOf((*Cache)(nil)).Elem(), options, opts...)
	if err != nil {
		return nil, errors.WithMessage(err, "refx.New failed")
	}
	return v.(Cache), nil
}

func GenerateCaptcha() string {
	buf := make([]byte, 8)
	_, _ = rand.Read(buf)
	return fmt.Sprintf("%06d", binary.LittleEndian.Uint64(buf)%1000000)
}

func GenerateToken() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}
