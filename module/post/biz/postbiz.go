package postbiz

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

type CreatePostStorage interface {
	Create(ctx context.Context, data *postmodel.Post) error
}

type createPostBiz struct {
	storage CreatePostStorage
}

func NewPostBiz(storage CreatePostStorage) *createPostBiz { return &createPostBiz{storage: storage} }

func (biz *createPostBiz) CreatePost(ctx context.Context, data *postmodel.Post) error {

	if err := biz.storage.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(postmodel.EntityName, err)
	}

	return nil
}
