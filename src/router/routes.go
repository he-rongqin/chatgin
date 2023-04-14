package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"org.chatgin/src/common"
	"org.chatgin/src/service"
)

func Router() *gin.Engine {

	r := gin.Default()
	// 全局中间件
	// Logger 中间件将日志写入 gin.DefaultWriter，即使你将 GIN_MODE 设置为 release。
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery 中间件会 recover 任何 panic。如果有 panic 的话，会写入 500。
	r.Use(gin.Recovery())

	r.POST("/auth/login", func(ctx *gin.Context) {
		var userLogin service.UserLoginForm
		// binding 校验
		if err := ctx.ShouldBind(&userLogin); err != nil {
			ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
			return
		}
		// 调用service 完成登录
		us := &service.UserService{}
		userInfo, err := us.Login(userLogin)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
			return
		}
		ctx.JSON(http.StatusBadRequest, common.ResData(userInfo))
	})
	return r
}
