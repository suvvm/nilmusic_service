package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/method"
	"suvvm.work/nilmusic_service/model"
)

// Login 登录接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func Login(ctx *gin.Context) {
	req := &model.LoginReq{}
	if err := ctx.BindJSON(req); err != nil {	// 读取请求
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
			"uid": 0,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoLogin(req))	// 执行登录
}

// Register 注册接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func Register(ctx *gin.Context) {
	req := &model.RegisterReq{}
	if err := ctx.BindJSON(req); err != nil {	// 读取请求
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoRegister(req))	// 执行注册
}
