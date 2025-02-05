package postbiz

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

type DeletePostStorage interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*postmodel.Post, error)
	Delete(ctx context.Context, id int) error
}

type deletePostBiz struct {
	storage   DeletePostStorage
	requester common.Requester
}

func NewDeletePostBiz(storage DeletePostStorage, requester common.Requester) *deletePostBiz {
	return &deletePostBiz{storage: storage, requester: requester}
}

func (biz *deletePostBiz) DeletePost(ctx context.Context, id int) error {

	oldData, err := biz.storage.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}

	if oldData.Status == "deleted" {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if oldData.AuthorID != biz.requester.GetUserId() {
		return common.ErrNotPermission(err)
	}

	if err := biz.storage.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(postmodel.EntityName, err)
	}

	return nil
}
