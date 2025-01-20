package middleware

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/tokenprovider/jwt"
	userstorage "Blog-CMS/module/user/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong auth header",
		"ErrWrongAuthHeader")
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	// Ensure the header has the correct format
	if len(parts) < 2 || parts[0] != "Bearer" || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(fmt.Errorf("invalid authorization header format"))
	}

	return parts[1], nil
}

// 1.Get token from header
// 2.Validate token and parse to payload
// 3.From the token payload, we use user_id to find from DB

func RequireAuth(appCtx appctx.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {

		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		storage := userstorage.NewSqlStorage(db)

		user, err := storage.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
