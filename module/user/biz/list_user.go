package userbiz

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

type ListUserRepo interface {
	ListDataWithCondition(
		context context.Context,
		paging *common.Paging,
		morekeys ...string,
	) ([]usermodel.User, error)
}

type listUserBiz struct {
	repo ListUserRepo
}

func NewListUserBiz(repo ListUserRepo) *listUserBiz {
	return &listUserBiz{repo: repo}
}

func (biz *listUserBiz) GetListUser(
	context context.Context,
	paging *common.Paging,
) ([]usermodel.User, error) {

	result, err := biz.repo.ListDataWithCondition(context, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
