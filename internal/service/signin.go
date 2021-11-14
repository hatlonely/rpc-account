package service

import (
	"context"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/storage"
)

func (s *Service) SignIn(ctx context.Context, req *api.SignInReq) (*api.SignInRes, error) {
	account, err := s.storage.GetAccountByPhoneOrEmail(ctx, req.Username)
	if err != nil {
		if err == storage.ErrInvalidUsername {
			return nil, rpcx.NewErrorf(errors.Errorf("user [%v] is invalid", req.Username), codes.InvalidArgument, "InvalidArgument", "user [%v] is invalid", req.Username)
		}
		return nil, errors.WithMessage(err, "storage.GetAccountByPhoneOrEmail failed")
	}

	if account.Password != req.Password {
		return nil, rpcx.NewErrorf(errors.New("password is incorrect"), codes.PermissionDenied, "Forbidden", "password is incorrect")
	}

	token, err := s.cache.SetToken(ctx, account)
	if err != nil {
		return nil, errors.WithMessagef(err, "cache.SetToken failed")
	}

	return &api.SignInRes{
		Token: token,
	}, nil
}
