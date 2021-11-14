package mysql

import (
	"context"
	"encoding/hex"

	"github.com/hatlonely/rpc-account/internal/storage"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

func (m *MySQL) PutAccount(ctx context.Context, article *storage.Account) (string, error) {
	article.ID = hex.EncodeToString(uuid.NewV4().Bytes())
	return article.ID, m.db.Create(ctx, article).Unwrap().Error
}

func (m *MySQL) GetAccount(ctx context.Context, id string) (*storage.Account, error) {
	var article storage.Account
	if err := m.db.
		Where(ctx, "`id`=?", id).
		First(ctx, &article).
		Unwrap().Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &article, nil
}

func (m *MySQL) UpdateAccount(ctx context.Context, article *storage.Account) error {
	return m.db.Model(ctx, article).Where(ctx, "`id`=?", article.ID).Updates(ctx, article).Unwrap().Error
}

func (m *MySQL) DelAccount(ctx context.Context, id string) error {
	return m.db.Delete(ctx, &storage.Account{ID: id}).Unwrap().Error
}
