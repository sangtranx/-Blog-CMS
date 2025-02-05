package poststorage

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
	"gorm.io/gorm"
)

func (s *sqlStorage) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	morekeys ...string,
) (*postmodel.Post, error) {

	var data postmodel.Post

	if err := s.db.Where(condition).First(&data).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &data, nil
}
