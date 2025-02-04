package poststorage

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

func (s *sqlStorage) Create(ctx context.Context, data *postmodel.Post) error {

	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
