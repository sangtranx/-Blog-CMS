package userstorage

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
	"time"
)

func (s *sqlStorage) FindUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*usermodel.User, error) {

	// create a timeout context
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	db := s.db.WithContext(timeoutCtx) //use context to set timeout for db query

	var user usermodel.User

	if err := db.Where(conditions).First(&user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user, nil
}
