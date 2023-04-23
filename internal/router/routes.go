package router

import (
	"github.com/gin-gonic/gin"
	"org.chatgin/internal/handler"
	"org.chatgin/internal/middleware"
)

func Router() *gin.Engine {

	router := gin.Default()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	router.Use(gin.Recovery())
	v1 := router.Group("api/v1")
	{
		v1.POST("/user/register", handler.UserRegisterEndpoint)
		v1.POST("/auth/login", handler.UserLoginEndpoint)
	}
	authorized := router.Group("api/v1").Use(middleware.AuthRequired())
	// 使用jwt 中间件鉴权
	// authorized.Use(middleware.AuthRequired())

	{
		authorized.GET("/user/:uid", handler.GetUserInfoEndpoint)
	}

	return router
}
