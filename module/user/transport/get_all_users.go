package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	usermodel "Blog-CMS/module/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
