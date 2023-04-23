package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"org.chatgin/internal/module"
	"org.chatgin/internal/service"
	common "org.chatgin/pkg/util"
)

// 用户注册
func UserRegisterEndpoint(ctx *gin.Context) {
	var userRegister module.UserRegisterForm
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
	var userLogin module.UserLoginForm
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
}

func GetUserInfoEndpoint(ctx *gin.Context) {
	id := ctx.Param("uid")
	userService := &service.UserService{}
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
		return
	}
	user, err := userService.GetInfo(uint(uid))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, common.ResError(http.StatusBadRequest, err))
		return
	}
	ctx.JSON(http.StatusBadRequest, common.ResData(user))

}
