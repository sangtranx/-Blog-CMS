package posttransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	postbiz "Blog-CMS/module/post/biz"
	poststorage "Blog-CMS/module/post/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// DeletePost godoc
// @Summary Delete a post with id
// @Description Delete a post with id
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id query string true "post id in base58 format"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /post/delete [post]
func DeletePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()
		idStr := c.Query("id")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		storage := poststorage.NewSqlStorage(db)
		biz := postbiz.NewDeletePostBiz(storage, requester)

		if err := biz.DeletePost(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("this post has been deleted successfully"))
	}
}
