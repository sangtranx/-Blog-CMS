package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	hasher2 "Blog-CMS/component/hasher"
	userbiz "Blog-CMS/module/user/biz"
	usermodel "Blog-CMS/module/user/model"
	userstorage "Blog-CMS/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Change password a user's password
// @Description Change a user's password with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param data body usermodel.UserChangePd true "User password changed"
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /change/password [post]
func UpdatePassword(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		u := c.MustGet(common.CurrentUser).(common.Requester)

		var req usermodel.UserChangePd

		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		storage := userstorage.NewSqlStorage(db)
		hasher := hasher2.NewSha256Hash()
		biz := userbiz.NewChangePasswordBiz(storage, hasher)

		if err := biz.ChangePassword(c.Request.Context(), u.GetUserId(), req.Password); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("Password updated successfully"))
	}
}
