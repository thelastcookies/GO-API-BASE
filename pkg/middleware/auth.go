package middleware

import (
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth authorize user
func Auth(paths ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ignore some path
		// eg: register, login, logout
		if len(paths) > 0 {
			path := c.Request.URL.Path
			pathsStr := strings.Join(paths, "|")
			reg := regexp.MustCompile("(" + pathsStr + ")")
			if reg.MatchString(path) {
				return
			}
		}

		// Parse the json web token.
		ctx, err := ParseRequest(c)
		if err != nil {
			//router.NewResponse().Error(c, errno.ErrInvalidToken)
			c.Abort()
			return
		}

		// set uid to context
		c.Set("uid", ctx.UserID)

		c.Next()
	}
}
