package chattransport

import (
	"Blog-CMS/component/appctx"
	chatbiz "Blog-CMS/module/chat/biz"
	chatmodel "Blog-CMS/module/chat/model"
	chatstorage "Blog-CMS/module/chat/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleDeepSeekQuery godoc
// @Summary Chat with deepseek ai
// @Description Chat with deepseek ai
// @Tags AI
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer <Token>"
// @Param data body chatmodel.DeepSeekRequest true "adding request"
// @Success 200 {object} common.SuccessRes
// @Failure 400 {object} common.AppError
// @Failure 401 {object} common.AppError
// @Failure 500 {object} common.AppError
// @Router /deepseek/chat [post]
func HandleDeepSeekQuery(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var Body chatmodel.DeepSeekRequest
		if err := c.ShouldBindJSON(&Body); err != nil {
			panic(err)
		}

		dbStorage := chatstorage.NewSqlStorage(appCtx.GetMainDBConnection())
		storageInstance := chatmodel.NewDeepSeekModel(Body.Model, Body.ApiKey, Body.ApiURL, Body.Messages)
		bizInstance := chatbiz.NewDeepSeekBiz(dbStorage)
		resp, err := bizInstance.GetDeepSeekResponse(*storageInstance)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, resp)
	}
}
