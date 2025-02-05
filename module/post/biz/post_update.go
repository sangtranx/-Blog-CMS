package postbiz

import (
	"Blog-CMS/common"
	postmodel "Blog-CMS/module/post/model"
	"context"
)

type PostUpdateStorage interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekeys ...string,
	) (*postmodel.Post, error)
	Update(ctx context.Context, data *postmodel.PostUpdate) error
}

type postUpdateBiz struct {
	storage PostUpdateStorage
	requester common.Requester
}

func NewPostUpdateBiz(storage PostUpdateStorage, requester common.Requester) *postUpdateBiz {
	return &postUpdateBiz{storage: storage, requester: requester}
}

func (biz *postUpdateBiz) PostUpdateBiz(ctx context.Context, data *postmodel.PostUpdate) error {

	oldData, err := biz.storage.FindDataWithCondition(ctx, map[string]interface{}{"id": data.ID})

	if err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	if oldData.Status == "deleted" {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, nil)
	}

	if oldData.AuthorID != biz.requester.GetUserId() {
		return common.ErrNotPermission(err)
	}

	if err := biz.storage.Update(ctx, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}