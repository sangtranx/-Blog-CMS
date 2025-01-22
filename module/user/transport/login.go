package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	hasher2 "Blog-CMS/component/hasher"
	"Blog-CMS/component/tokenprovider/jwt"
	userbiz "Blog-CMS/module/user/biz"
	usermodel "Blog-CMS/module/user/model"
	userstorage "Blog-CMS/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary Authenticate user
// @Description Authenticate user with email and password
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]interface{} "Successfully authenticated"
// @Failure 400 {object} map[string]interface{} "Invalid input"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /authenticate [post]
func Login(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data usermodel.UserLogin

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		storage := userstorage.NewSqlStorage(db)
		tokenProvider := jwt.NewJWTProvider(appCtx.SecretKey())
		hasher := hasher2.NewSha256Hash()

		biz := userbiz.NewLoginBusiness(storage, tokenProvider, hasher, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), data)

		if err != nil {
			common.ErrInvalidRequest(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
