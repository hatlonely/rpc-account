package service

import (
	"context"
	"strings"

	"github.com/go-redis/redis"
	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/storage"
)

func (s *AccountService) SignUp(ctx context.Context, req *api.SignUpReq) (*empty.Empty, error) {
	key := "captcha_" + req.Email
	val, err := s.redisCli.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, rpcx.NewErrorf(errors.New("captcha is not exists"), codes.InvalidArgument, "InvalidArgument", "captcha is not exists")
	}
	if err != nil {
		return nil, errors.Wrapf(err, "redis get key [%v] failed", key)
	}
	if req.Captcha != val {
		return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "captcha is not match")
	}

	birthday, err := cast.ToTimeE(req.Birthday)
	if err != nil {
		return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "invalid birthday format")
	}

	user := &storage.Account{
		Email:    req.Email,
		Phone:    req.Phone,
		Name:     req.Name,
		Password: req.Password,
		Birthday: birthday,
		Gender:   int(req.Gender),
		Avatar:   req.Avatar,
	}
	if err := s.mysqlCli.Create(ctx, user).Unwrap().Error; err != nil {
		switch err.(type) {
		case *mysql.MySQLError:
			e := err.(*mysql.MySQLError)
			if e.Number == 1062 {
				if strings.Contains(e.Message, "accounts.email_idx") {
					return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "account [%v] exists", req.Email)
				}
				if strings.Contains(e.Message, "accounts.phone_idx") {
					return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "account [%v] exists", req.Phone)
				}
				if strings.Contains(e.Message, "accounts.name_idx") {
					return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "account [%v] exists", req.Name)
				}
			}
		}
		return nil, errors.Wrap(err, "mysql create failed")
	}

	return &empty.Empty{}, nil
}
