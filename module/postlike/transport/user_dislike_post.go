package postliketranspot

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	postlikebiz "Blog-CMS/module/postlike/biz"
	postlikestorage "Blog-CMS/module/postlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserDislikePost godoc
// @Summary Dislike a post with id
// @Description Dislike a post with id
// @Tags posts
// @Accept  json
// @Produce  json
// @Param id query string true "post id in base58 format"
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /post/dislike [delete]
func UserDislikePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()
		idStr := c.Query("id")

		postId, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		storage := postlikestorage.NewSQLStorage(db)
		biz := postlikebiz.NewUserDisLikePost(storage, appCtx.GetPubsub())

		if err := biz.UserDisLikePost(c.Request.Context(), requester.GetUserId(), postId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
