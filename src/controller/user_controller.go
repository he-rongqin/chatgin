package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"org.chatgin/src/common"
	"org.chatgin/src/service"
)

// 用户注册
func UserRegisterEndpoint(ctx *gin.Context) {
	var userRegister service.UserRegisterForm
	if err := ctx.ShouldBind(&userRegister); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
		return
	}
	// 调用service 完成注册
	us := &service.UserService{}
	err := us.Register(userRegister)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(http.StatusBadRequest, common.ResSuccess())
}

func UserLoginEndpoint(ctx *gin.Context) {
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
	// 组装 map，因为userinfo 属性是私有的
	userJson := map[string]any{
		"id":       userInfo.GetId(),
		"username": userInfo.GetUsername(),
		"state":    userInfo.GetState(),
		"token":    userInfo.GetToken(),
	}
	ctx.JSON(http.StatusBadRequest, common.ResData(userJson))
}