package postlikestorage

import (
	"Blog-CMS/common"
	postlikemodel "Blog-CMS/module/postlike/model"
	"context"
)

func (s *sqlstorage) Create(ctx context.Context, data *postlikemodel.PostLike) error {

	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
