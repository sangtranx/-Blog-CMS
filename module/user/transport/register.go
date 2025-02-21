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

// Register godoc
// @Summary Register a new user
// @Description Register a new user with the provided information
// @Tags users
// @Accept  json
// @Produce  json
// @Param data body usermodel.UserRegister true "User registration data"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /register [post]
func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		// validate data
		if err := data.Validate(); err != nil {
			panic(err)
		}

		storage := userstorage.NewSqlStorage(db)
		hasher := hasher2.NewSha256Hash()
		biz := userbiz.NewRegisterUserBusiness(storage, hasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
