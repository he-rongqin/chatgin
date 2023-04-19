package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"org.chatgin/src/common"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		tokenStr := ctx.GetHeader(common.HEADER_AUTHORIZATION_KEY)
		if tokenStr == "" {
			ctx.JSON(common.NOT_AUTHENTICATION, common.ResErrorMsg(common.NOT_AUTHENTICATION, "无权限访问该资源，请登录之后再试"))
			// 拒绝请求
			ctx.Abort()
			return
		}
		if !strings.HasPrefix(tokenStr, common.TOKEN_PREFIX) {
			ctx.JSON(http.StatusBadRequest, common.ResErrorMsg(http.StatusBadRequest, "无权限访问该资源，请登录之后再试"))
			// 拒绝请求
			ctx.Abort()
			return
		}
		// 解析token
		token := &Token{}
		claims, err := token.Analyze(strings.Replace(tokenStr, "Bearer ", "", -1))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, common.ResErrorMsg(http.StatusBadRequest, "非法请求，token解析失败"))
			// 拒绝请求
			ctx.Abort()
			return
		}
		// 令牌过期
		if time.Now().Unix() > claims.ExpiresAt {
			ctx.JSON(http.StatusBadRequest, common.ResErrorMsg(http.StatusBadRequest, "会话已超时，清重新登录"))

			// 拒绝请求
			ctx.Abort()
			return
		}
		// todo 刷新token

		ctx.Next()

	}
}
