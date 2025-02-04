package posttransport

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	postbiz "Blog-CMS/module/post/biz"
	postmodel "Blog-CMS/module/post/model"
	poststorage "Blog-CMS/module/post/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateNewPost godoc
// @Summary create a new post
// @Description create a new post
// @Tags posts
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Param data body postmodel.PostCreate true "post information"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /post/create [post]
func CreateNewPost(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()

		var data postmodel.Post

		if err := c.BindJSON(&data); err != nil {
			panic(err)
		}

		storage := poststorage.NewSqlStorage(db)
		biz := postbiz.NewPostBiz(storage)

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
