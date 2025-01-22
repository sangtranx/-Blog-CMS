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
// @Success 200 {object} map[string]interface{} "Successfully retrieved profile"
// @Failure 401 {object} map[string]interface{} "Unauthorized"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /profile [get]
func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		u := c.MustGet(common.CurrentUser).(common.Requester)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}
