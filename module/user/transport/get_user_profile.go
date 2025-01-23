package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	usermodel "Blog-CMS/module/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserProfile godoc
// @Summary Get user profile by ID
// @Description Get the profile of a user by their ID
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id query string true "User ID in base58 format"
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /admin/userProfile [get]
func GetUserProfile(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		v := c.Query("id")

		uid, err := common.FromBase58(v)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})

			return
		}

		var data usermodel.User

		db.Where("id = ?", uid.GetLocalID()).First(&data)
		data.Mask(data.GetRole() == common.AdminRole)

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
