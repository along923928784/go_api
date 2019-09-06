package middleware

import (
	"jiyue.im/pkg/errno"
	"jiyue.im/pkg/token"
	"jiyue.im/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(level int32) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		ctx, err := token.ParseRequest(c)
		// if err != nil || ctx.Scope < level {
		// 	service.SendResponse(c, errno.ErrTokenInvalid, nil)
		// 	c.Abort()
		// 	return
		// }
		if err != nil {
			service.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		if ctx.Scope < level {
			service.SendResponse(c, errno.ErrForbbiden, nil)
			c.Abort()
			return
		}

		c.Set("userContext", ctx)
		c.Next()
	}
}
