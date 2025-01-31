package userbiz

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

type ChangePasswordStorage interface {
	UpdatePassword(ctx context.Context, userID int, hashedPassword, salt string) error
}

type changePasswordBiz struct {
	storage ChangePasswordStorage
	hasher  Hasher
}

func NewChangePasswordBiz(storage ChangePasswordStorage, hasher Hasher) *changePasswordBiz {
	return &changePasswordBiz{storage: storage, hasher: hasher}
}

func (biz *changePasswordBiz) ChangePassword(ctx context.Context, userID int, newPassword string) error {

	salt := common.GenSalt(50)
	password := biz.hasher.Hash(newPassword + salt)

	if err := biz.storage.UpdatePassword(ctx, userID, password, salt); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
