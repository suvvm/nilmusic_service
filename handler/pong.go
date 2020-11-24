package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"suvvm.work/nilmusic_service/common"
)

// Pong ping-pong测试接口
//
// 入参
//	ctx *gin.Context	// 上下文参数
func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": common.HandlerSuccess,
		"msg": "pong",
	})
}
