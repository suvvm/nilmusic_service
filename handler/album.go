package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/method"
	"suvvm.work/nilmusic_service/model"
)

// AllAlbum 获取当前用户全部专辑接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func AllAlbum (ctx *gin.Context) {
	uid, err := strconv.Atoi(ctx.DefaultQuery("uid", "0"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadPathErr,
			"msg": common.HandlerReadPathErrMsg,
			"album_list": make([]*model.Album, 0),
		})
	}
	ctx.JSON(http.StatusOK, method.DoGetAllAlbum(uid))
}

// CreateAlbum 创建新专辑接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func CreateAlbum (ctx *gin.Context) {
	req := &model.CreateAlbumReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
			"aid": 0,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoCreateAlbum(req))
}

// DeleteAlbum 删除专辑接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func DeleteAlbum(ctx *gin.Context) {
	req := &model.DeleteAlbumReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoDeleteAlbum(req))
}
