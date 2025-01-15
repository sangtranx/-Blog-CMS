package usertransport

import (
	"Blog-CMS/component/appctx"
	usermodel "Blog-CMS/module/user/model"
	userstorage "Blog-CMS/module/user/storage"
	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		storage := userstorage.NewSqlStorage(db)

		//biz := userbiz.NewRegisterUserBusiness(storage)
	}
}
