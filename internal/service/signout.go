package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-account/api/gen/go/api"
)

func (s *AccountService) SignOut(ctx context.Context, req *api.SignOutReq) (*empty.Empty, error) {
	if err := s.kv.Del(req.Token); err != nil {
		return nil, errors.Wrapf(err, "kv del [%v] failed", req.Token)
	}

	return &empty.Empty{}, nil
}
