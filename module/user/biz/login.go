package userbiz

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/tokenprovider"
	usermodel "Blog-CMS/module/user/model"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*usermodel.User, error)
}

type LoginBusiness struct {
	storage       LoginStorage
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storage LoginStorage, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *LoginBusiness {
	return &LoginBusiness{
		storage:       storage,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

// 1.Find user, email
// 2.Hash pass from input and compare with pass in db
// 3.Provider: issue JWT token for client
// 4. Return token(s)

func (business *LoginBusiness) Login(
	ctx context.Context,
	appCtx appctx.AppContext,
	data usermodel.UserLogin,
) (*tokenprovider.Token, error) {

	user, err := business.storage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailnameOrPasswordInvalid
	}

	pwd := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != pwd {
		// register password fail
		data.RegisterFailedAttempt(appCtx)
		return nil, usermodel.ErrEmailnameOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	// delete failed attempt
	data.ResetAttempts(appCtx)

	return accessToken, nil
}
