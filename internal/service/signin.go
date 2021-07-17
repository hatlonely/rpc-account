package service

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/go-kit/strx"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/storage"
)

func GenerateToken() string {
	return hex.EncodeToString(uuid.NewV4().Bytes())
}

func (s *AccountService) SignIn(ctx context.Context, req *api.SignInReq) (*api.SignInRes, error) {
	a := &storage.Account{}
	if strx.RePhone.MatchString(req.Username) {
		if err := s.mysqlCli.Where("phone=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, rpcx.NewErrorf(err, codes.PermissionDenied, "Forbidden", "user [%v] not exist", req.Username)
			}
			return nil, errors.Wrapf(err, "mysql select user [%v] failed", req.Username)
		}
	} else if strx.ReEmail.MatchString(req.Username) {
		if err := s.mysqlCli.Where("email=?", req.Username).First(a).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, rpcx.NewErrorf(err, codes.PermissionDenied, "Forbidden", "user [%v] not exist", req.Username)
			}
			return nil, errors.Wrapf(err, "mysql select user [%v] failed", req.Username)
		}
	} else {
		return nil, rpcx.NewErrorf(errors.Errorf("user [%v] is invalid", req.Username), codes.InvalidArgument, "InvalidArgument", "user [%v] is invalid", req.Username)
	}

	if a.Password != req.Password {
		return nil, rpcx.NewErrorf(errors.New("password is incorrect"), codes.PermissionDenied, "Forbidden", "password is incorrect")
	}

	token := GenerateToken()

	if err := s.kv.Set(token, a); err != nil {
		return nil, errors.Wrapf(err, "redis set token [%v] failed", token)
	}

	return &api.SignInRes{
		Token: token,
	}, nil
}
