package service

import (
	"context"
	"time"

	"github.com/hatlonely/go-kit/cast"
	"github.com/hatlonely/go-kit/rpcx"
	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/hatlonely/rpc-account/internal/storage"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

func (s *Service) PutAccount(ctx context.Context, account *api.Account) (*api.AccountID, error) {
	birthday, err := cast.ToTimeE(account.Birthday)
	if err != nil {
		return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "invalid birthday format")
	}

	id, err := s.storage.PutAccount(ctx, &storage.Account{
		Email:    account.Email,
		Phone:    account.Phone,
		Name:     account.Name,
		Password: account.Password,
		Birthday: birthday,
		Gender:   int(account.Gender),
		Avatar:   account.Avatar,
	})
	if err != nil {
		return nil, errors.WithMessagef(err, "storage.PutAccount failed")
	}

	return &api.AccountID{
		Id: id,
	}, nil
}

func (s *Service) UpdateAccount(ctx context.Context, account *api.Account) (*api.Empty, error) {
	var birthday time.Time
	if account.Birthday != "" {
		var err error
		birthday, err = cast.ToTimeE(account.Birthday)
		if err != nil {
			return nil, rpcx.NewErrorf(err, codes.InvalidArgument, "InvalidArgument", "invalid birthday format")
		}
	}

	err := s.storage.UpdateAccount(ctx, &storage.Account{
		Email:    account.Email,
		Phone:    account.Phone,
		Name:     account.Name,
		Password: account.Password,
		Birthday: birthday,
		Gender:   int(account.Gender),
		Avatar:   account.Avatar,
	})
	if err != nil {
		return nil, errors.WithMessagef(err, "storage.UpdateAccount failed")
	}

	return &api.Empty{}, nil
}

func (s *Service) GetAccount(ctx context.Context, id *api.AccountID) (*api.Account, error) {
	account, err := s.storage.GetAccount(ctx, id.Id)
	if err != nil {
		return nil, errors.WithMessage(err, "storage.GetAccount failed")
	}

	return &api.Account{
		Id:       account.ID,
		Email:    account.Email,
		Phone:    account.Phone,
		Name:     account.Name,
		Password: account.Password,
		Birthday: account.Birthday.Format(time.RFC3339),
		Gender:   api.Gender(account.Gender),
		Avatar:   account.Avatar,
	}, nil
}

func (s *Service) DelAccount(ctx context.Context, id *api.AccountID) (*api.Empty, error) {
	if err := s.storage.DelAccount(ctx, id.Id); err != nil {
		return nil, errors.WithMessage(err, "storage.DelAccount failed")
	}
	return &api.Empty{}, nil
}
