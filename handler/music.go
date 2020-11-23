package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/method"
	"suvvm.work/nilmusic_service/model"
)

// AllMusic 获取目标专辑全部音乐接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func AllMusic(ctx *gin.Context) {
	aid, err := strconv.Atoi(ctx.DefaultQuery("aid", "0"))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadPathErr,
			"msg": common.HandlerReadPathErrMsg,
			"music_list": make([]*model.Music, 0),
		})
	}
	ctx.JSON(http.StatusOK, method.DoGetAllMusic(aid))
}

// AddMusic 添加音乐至目标专辑接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func AddMusic(ctx *gin.Context) {
	req := &model.AddMusicReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoAddMusic(req))
}

// MdfMusic 修改音乐信息接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func MdfMusic(ctx *gin.Context) {
	req := &model.MdfMusicReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoMdfMusic(req))
}

// DelMusic 在目标专辑中删除音乐接口处理方法
//
// 入参
//	ctx *gin.Context	// 上下文参数
func DelMusic(ctx *gin.Context) {
	req := &model.DeleteMusicReq{}
	if err := ctx.BindJSON(req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": common.HandlerReadBodyErr,
			"msg": common.HandlerReadBodyErrMsg,
		})
		return
	}
	ctx.JSON(http.StatusOK, method.DoDeleteMusic(req))
}
