package userstorage

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

func (s *sqlStorage) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {

	db := s.db

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
