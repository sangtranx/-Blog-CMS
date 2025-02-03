package middleware

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"errors"
	"github.com/gin-gonic/gin"
)

func RoleRequired(appCtx appctx.AppContext, allowRole ...string) func(c *gin.Context) {

	return func(c *gin.Context) {

		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRole {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNotPermission(errors.New("Invalid role user")))
		}

		c.Next()
	}
}
