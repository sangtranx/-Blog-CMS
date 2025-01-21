package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	usermodel "Blog-CMS/module/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
