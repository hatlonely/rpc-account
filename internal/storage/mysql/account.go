package mysql

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/go-kit/strx"
	"github.com/hatlonely/rpc-account/internal/storage"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (m *MySQL) PutAccount(ctx context.Context, account *storage.Account) (string, error) {
	account.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return account.ID, m.db.Create(ctx, account).Unwrap().Error
}

func (m *MySQL) GetAccount(ctx context.Context, id string) (*storage.Account, error) {
	var account storage.Account
	if err := m.db.
		Where(ctx, "`id`=?", id).
		First(ctx, &account).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &account, nil
}

func (m *MySQL) UpdateAccount(ctx context.Context, account *storage.Account) error {
	return m.db.Model(ctx, account).Where(ctx, "`id`=?", account.ID).Updates(ctx, account).Unwrap().Error
}

func (m *MySQL) DelAccount(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Account{ID: id}).Unwrap().Error
}

func (m *MySQL) GetAccountByPhoneOrEmail(ctx context.Context, username string) (*storage.Account, error) {
	var account storage.Account
	if strx.RePhone.MatchString(username) {
		if err := m.db.
			Where(ctx, "`phone`=?", username).
			First(ctx, &account).
			Unwrap().Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}

			return nil, err
		}

		return &account, nil
	}

	if strx.ReEmail.MatchString(username) {
		if err := m.db.
			Where(ctx, "`email`=?", username).
			First(ctx, &account).
			Unwrap().Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return nil, nil
			}

			return nil, err
		}

		return &account, nil
	}

	return nil, storage.ErrInvalidUsername
}
