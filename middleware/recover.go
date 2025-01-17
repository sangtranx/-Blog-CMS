package middleware

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"github.com/gin-gonic/gin"
)

func Recover(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {

				c.Header("Content-Type", "apllication/json")

				if appErr, ok := err.(*common.AppError); ok {

					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					panic(err) // to upstack
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err) // to upstack
				return
			}
		}()

		c.Next()
	}
}
