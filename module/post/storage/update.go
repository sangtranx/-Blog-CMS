package poststorage

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

func (s *sqlStorage) Update(ctx context.Context, data *postmodel.PostUpdate) error {

	if err := s.db.
		Table(postmodel.Post{}.TableName()).
		Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
