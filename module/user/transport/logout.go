package usertransport

import (
	"Blog-CMS/component/appctx"
	cache "Blog-CMS/component/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @Summary log out
// @Description log out user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /logout [post]
func Logout(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		expiry := 100
		err := cache.AddToBlackList(appCtx.GetRedisDBConnection(), token, expiry)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not revoke token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
	}
}
