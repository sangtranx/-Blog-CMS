package poststorage

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
	"gorm.io/gorm"
	"log"
)

func (s *sqlStorage) Update(ctx context.Context, data *postmodel.PostUpdate) error {

	if err := s.db.
		Table(postmodel.Post{}.TableName()).
		Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}

func (s *sqlStorage) IncreaseLikeCount(ctx context.Context, postId int) error {

	db := s.db.Table(postmodel.Post{}.TableName())

	if err := db.Where("id = ?", postId).Update("likes", gorm.Expr("likes + ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	log.Println("Increase likes count ", postId)
	return nil
}

func (s *sqlStorage) DecreaseLikeCount(ctx context.Context, postId int) error {

	db := s.db.Table(postmodel.Post{}.TableName())

	if err := db.Where("id = ?", postId).Update("likes", gorm.Expr("likes - ?", 1)).Error; err != nil {
		return common.ErrDB(err)
	}

	log.Println("Decrease likes count ", postId)
	return nil
}
