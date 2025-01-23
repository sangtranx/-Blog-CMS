package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Profile godoc
// @Summary Get user profile
// @Description Get the profile of the authenticated user
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /profile [get]
func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		u := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
