package usertransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	userbiz "Blog-CMS/module/user/biz"
	userstorage "Blog-CMS/module/user/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(appCtx appctx.AppContext) gin.HandlerFunc {

	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var papingData common.Paging

		if err := c.ShouldBindQuery(&papingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		papingData.FullFill()

		storage := userstorage.NewSqlStorage(db)
		biz := userbiz.NewListUserBiz(storage)

		result, err := biz.GetListUser(c.Request.Context(), &papingData)

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": common.NewErrorResponse(err, "can not get list user", "StatusInternalServerError", ""),
			})
		}

		for i := range result {
			result[i].Mask(result[i].GetRole() == common.CurrentUser)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, papingData, nil))
	}
}
