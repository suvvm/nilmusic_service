package main

import (
	"github.com/gin-gonic/gin"
	"suvvm.work/nilmusic_service/handler"
)

func register(r *gin.Engine) {
	// 用户管理
	r.POST("/nilmusic/user/login", handler.Login)
	r.POST("/nilmusic/user/register", handler.Register)
	// 歌单管理
	r.GET("/nilmusic/list/all", handler.AllList)				// 获取当前用户全部歌单
	r.POST("/nilmusic/list/create", handler.CreateList)		// 创建歌单
	r.DELETE("/nilmusic/list/delete", handler.DeleteList)		// 删除歌单
	// 歌单音乐操作
	r.GET("/nilmusic/list/music", handler.AllMusic)			// 获取歌单中全部音乐
	r.POST("/nilmusic/list/music/add", handler.AddMusic)		// 向歌单中添加音乐
	r.PUT("/nilmusic/list/music/mdf", handler.MdfMusic)		// 修改歌单中的音乐
	r.DELETE("/nilmusic/list/music/delete", handler.DelMusic)	// 删除歌单中的音乐
}
