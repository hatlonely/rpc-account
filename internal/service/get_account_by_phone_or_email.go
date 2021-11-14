package service

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
	"github.com/pkg/errors"
)

func (s *Service) GetAccountByPhoneOrEmail(ctx context.Context, req *api.GetAccountByPhoneOrEmailReq) (*api.Account, error) {
	account, err := s.storage.GetAccountByPhoneOrEmail(ctx, req.Username)
	if err != nil {
		return nil, errors.WithMessage(err, "storage.GetAccountByPhoneOrEmail failed")
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
