package postlikestorage

import (
	"Blog-CMS/common"
	postlikemodel "Blog-CMS/module/postlike/model"
	"context"
)

func (s *sqlstorage) Delete(ctx context.Context, UserId, postId int) error {

	db := s.db

	if err := db.Table(postlikemodel.PostLike{}.TableName()).
		Where("user_id = ? and post_id = ?", UserId, postId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
