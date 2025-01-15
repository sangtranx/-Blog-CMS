package userbiz

import (
	"Blog-CMS/common"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, condition map[string]interface{}, moreInfos ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	storage RegisterStorage
	hasher  Hasher
}

func NewRegisterUserBusiness(storage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{storage: storage, hasher: hasher}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {

	user, _ := business.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	Salt := common.GenSalt(50)
	data.Password = business.hasher.Hash(data.Password + Salt)
	data.Role = "User"

	if err := business.storage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
