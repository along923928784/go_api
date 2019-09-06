package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"jiyue.im/pkg/auth"
	"jiyue.im/server/middleware"
	"jiyue.im/service"
)

func LoadRouter(mw ...gin.HandlerFunc) *gin.Engine {
	g := gin.New()

	if gin.Mode() == "debug" {
		g.Use(gin.Logger())
	}
	g.Use(gin.Recovery())
	g.Use(mw...)

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"你访问的页面不存在": "XXXXXXXXXXXXXXXXX",
		})
	})

	ping(g)

	u := g.Group("/v1/user")
	{
		u.POST("/register", service.UserRegister)
		u.POST("/token", service.GetToken)
	}

	g.POST("/v1/token/verify", middleware.AuthMiddleware(auth.USER), service.VerifyToken)

	classic := g.Group("/v1/classic")
	classic.Use(middleware.AuthMiddleware(auth.USER))
	{
		classic.GET("/latest", service.GetLatest)
	}
	// g.POST("/v1/token/verify1", middleware.AuthMiddleware(9), service.VerifyToken)
	// u := g.Group("/v1/user")
	// u.Use(middleware.AuthMiddleware())
	// {
	// 	u.POST("", user.Create)
	// 	u.DELETE("/:id", user.Delete)
	// 	u.PUT("/:id", user.Update)
	// 	u.GET("", user.List)
	// 	u.GET("/:username", user.Get)
	// }

	// // The user handlers, requiring authentication
	// u := g.Group("/v1/user")
	// u.Use(middleware.AuthMiddleware())
	// {
	// 	u.POST("", user.Create)
	// 	u.DELETE("/:id", user.Delete)
	// 	u.PUT("/:id", user.Update)
	// 	u.GET("", user.List)
	// 	u.GET("/:username", user.Get)
	// }

	// // The health check handlers
	// svcd := g.Group("/sd")
	// {
	// 	svcd.GET("/health", sd.HealthCheck)
	// 	svcd.GET("/disk", sd.DiskCheck)
	// 	svcd.GET("/cpu", sd.CPUCheck)
	// 	svcd.GET("/ram", sd.RAMCheck)
	// }

	return g
}

func ping(g *gin.Engine) {
	g.GET("/ping", service.HealthCheck)
}
