package userstorage

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

func (s *sqlStorage) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*usermodel.User, error) {

	db := s.db

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
