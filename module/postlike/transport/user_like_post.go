package postliketranspot

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	postlikebiz "Blog-CMS/module/postlike/biz"
	postlikemodel "Blog-CMS/module/postlike/model"
	postlikestorage "Blog-CMS/module/postlike/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserlikePost godoc
// @Summary like a new post
// @Description like a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Param id query string true "post id in base58 format"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /post/like [post]
func UserlikePost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()
		idStr := c.Query("id")

		postId, err := strconv.Atoi(idStr)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		data := postlikemodel.PostLike{
			PostId: postId,
			UserId: requester.GetUserId(),
		}

		storage := postlikestorage.NewSQLStorage(db)
		biz := postlikebiz.NewUserLikePost(storage, appCtx.GetPubsub())

		if err := biz.UserLikePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
