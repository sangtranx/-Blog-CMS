package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	usermodel "Blog-CMS/module/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllUser godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /admin/users [get]
func GetAllUser(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var datas []usermodel.User

		db.Order("id desc").Find(&datas)

		for i := range datas {
			datas[i].Mask(datas[i].GetRole() == common.AdminRole) // Gọi Mask
		}

		c.JSON(http.StatusOK, gin.H{
			"datas": datas,
		})
	}
}
