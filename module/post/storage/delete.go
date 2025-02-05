package poststorage

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

func (s *sqlStorage) Delete(ctx context.Context, id int) error {

	if err := s.db.
		Table(postmodel.Post{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{"status": "deleted"}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
