package posttransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	postbiz "Blog-CMS/module/post/biz"
	postmodel "Blog-CMS/module/post/model"
	poststorage "Blog-CMS/module/post/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdatePost godoc
// @Summary create a new post
// @Description create a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Param data body postmodel.PostUpdate true "post information"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /post/update [post]
func UpdatePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()

		var data postmodel.PostUpdate

		if err := c.BindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		storage := poststorage.NewSqlStorage(db)
		biz := postbiz.NewPostUpdateBiz(storage, requester)

		if err := biz.PostUpdateBiz(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("this post has been updated successfully"))
	}
}
