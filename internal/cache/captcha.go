package cache

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"reflect"

	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"
)

type Cache interface {
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
