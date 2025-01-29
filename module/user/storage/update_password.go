package userstorage

import (
	usermodel "Blog-CMS/module/user/model"
	"context"
)

func (s *sqlStorage) UpdatePassword(ctx context.Context, userID int, hashedPassword, salt string) error {

	db := s.db

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id = ?", userID).
		Updates(map[string]interface{}{
			"password": hashedPassword,
			"salt":     salt,
		}).Error; err != nil {
		return err
	}

	return nil
}
