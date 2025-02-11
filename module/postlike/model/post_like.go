package postlikemodel

import (
	"Blog-CMS/common"
	"fmt"
	"time"
)

type PostLike struct {
	PostId    int                `json:"post_id" gorm:"column:post_id"`
	UserId    int                `json:"user_id" gorm:"column:user_id"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false"`
	CreatedAt time.Time          `json:"created_at" gorm:"column:created_at"`
}

func (PostLike) TableName() string {
	return "post_like"
}

func (l *PostLike) GetPostId() int {
	return l.PostId
}

func ErrCannotLikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this Post"),
		fmt.Sprintf("ErrCannotLikePost"),
	)
}

func ErrCannotDisLikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot Dislike this Post"),
		fmt.Sprintf("ErrCannotDisLikePost"),
	)
}
