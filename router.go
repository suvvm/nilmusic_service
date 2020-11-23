package main

import (
	"github.com/gin-gonic/gin"
	"suvvm.work/nilmusic_service/handler"
)

func register(r *gin.Engine) {
	// 用户管理
	r.POST("/nilmusic/user/login", handler.Login)
	r.POST("/nilmusic/user/register", handler.Register)
	// 专辑管理
	r.GET("/nilmusic/album/all", handler.AllAlbum)				// 获取当前用户全部专辑
	r.POST("/nilmusic/album/create", handler.CreateAlbum)			// 创建专辑
	r.DELETE("/nilmusic/album/delete", handler.DeleteAlbum)		// 删除专辑
	// 专辑音乐操作
	r.GET("/nilmusic/album/music", handler.AllMusic)				// 获取专辑中全部音乐
	r.POST("/nilmusic/album/music/add", handler.AddMusic)			// 向专辑中添加音乐
	r.PUT("/nilmusic/album/music/mdf", handler.MdfMusic)			// 修改专辑中的音乐
	r.DELETE("/nilmusic/album/music/delete", handler.DelMusic)	// 删除专辑中的音乐
}
