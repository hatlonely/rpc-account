package service

import (
	"context"

	"github.com/hatlonely/rpc-account/internal/cache"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/storage"
)

func (s *AccountService) SignUp(ctx context.Context, req *api.SignUpReq) (*empty.Empty, error) {
	val, err := s.cache.GetCaptcha(ctx, req.Email)
	if err == cache.ErrNotFound {
		return nil, rpcx.NewErrorf(errors.New("captcha is not exists"), codes.InvalidArgument, "InvalidArgument", "captcha is not exists")
	}
	if err != nil {
		return nil, errors.WithMessage(err, "cache.GetCaptcha failed")
	}
	if req.Captcha != val {
		return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "captcha is not match")
	}

	birthday, err := cast.ToTimeE(req.Birthday)
	if err != nil {
		return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "invalid birthday format")
	}

	_, err = s.storage.PutAccount(ctx, &storage.Account{
		Email:    req.Email,
		Phone:    req.Phone,
		Name:     req.Name,
		Password: req.Password,
		Birthday: birthday,
		Gender:   int(req.Gender),
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, errors.Wrap(err, "s.storage.PutAccount failed")
	}

	return &empty.Empty{}, nil
}
