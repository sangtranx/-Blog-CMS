package middleware

import (
	"Blog-CMS/common"
	"Blog-CMS/component/appctx"
	"Blog-CMS/component/tokenprovider/jwt"
	userstorage "Blog-CMS/module/user/storage"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"wrong auth header",
		"ErrWrongAuthHeader")
}

func extractTokenFromHeaderString(s string) (string, error) {
	// Trim any leading or trailing whitespace
	s = strings.TrimSpace(s)

	// Check if the header starts with "Bearer "
	if strings.HasPrefix(s, "Bearer ") {
		// Split the header into parts
		parts := strings.Split(s, " ")
		if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
			return "", ErrWrongAuthHeader(fmt.Errorf("invalid authorization header format"))
		}
		return parts[1], nil
	}

	// If the header does not start with "Bearer ", assume it's just the token
	if s == "" {
		return "", ErrWrongAuthHeader(fmt.Errorf("empty authorization header"))
	}

	return s, nil
}

// 1.Get token from header
// 2.Validate token and parse to payload
// 3.From the token payload, we use user_id to find from DB

func RequireAuth(appCtx appctx.AppContext) func(c *gin.Context) {

	tokenProvider := jwt.NewJWTProvider(appCtx.SecretKey(), appCtx.GetRedisDBConnection())

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

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
