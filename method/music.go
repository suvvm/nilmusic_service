package method

import (
	"fmt"
	"log"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/dal/db"
	"suvvm.work/nilmusic_service/model"
)

// DoGetAllMusic 获取专辑全部音乐业务逻辑
// 判断专辑aid是否获取成功，并根据aid获取对应音乐列表，并构造获取音乐响应
//
// 入参
//	aid int				// 专辑aid
// 返回
//	*model.AllMusicResp	// 获取音乐响应
func DoGetAllMusic (aid int) *model.AllMusicResp {
	resp := &model.AllMusicResp{
		Code: common.HandlerSuccess,
		Msg: "get music success",
		MusicList: make([]*model.Music, 0),
	}
	if aid == 0 {
		resp.Code = common.HandlerReadPathErr
		resp.Msg = common.HandlerReadPathErrMsg
		return resp
	}
	albumList, err := GetMusicByAID(aid)	// 根据aid获取音乐列表
	if err != nil {
		resp.Code = common.HandlerDBSelectErr
		resp.Msg = fmt.Sprintf("db get music list err=%v", err)
		return resp
	}
	resp.MusicList = albumList
	return resp
}

// GetMusicByAID 根据专辑ID获取目标专辑的全部音乐
// 根据专辑ID查询到专辑MID列表后，遍历MID获取音乐信息并返回音乐列表
//
// 入参
//	aid int			// 音乐aid
// 返回
//	[]*model.Music	// 歌曲列表
//	error			// 错误信息
func GetMusicByAID (aid int)  ([]*model.Music, error) {
	musicList := make([]*model.Music, 0)
	albumMusic := &model.AlbumMusic{Aid: aid}
	albumMusicList, err := db.GetAlbumMusic(albumMusic)	// album_music获取aid对应的全部关系
	if err != nil {
		return nil, err
	}
	for _, albumMusic := range *albumMusicList {	// 遍历关系，根据mid获取音乐详细信息
		music := &model.Music{ID: albumMusic.Mid}
		musics, err := db.GetMusic(music)
		if err != nil {
			log.Printf("GetMusicByAID: get music by MID=%d err=%v", albumMusic.Mid, err)
			continue
		}
		for index := range *musics {	// 追加专辑至musicList
			musicList = append(musicList, &(*musics)[index])
		}
	}
	return musicList, nil
}

// DoAddMusic 新增音乐
// 创建音乐并添加目标专辑与新建音乐的关系
//
// 入参
//	req *model.AddMusicReq	// 新增音乐请求
// 返回
//	*model.AddMusicResp		// 新增音乐响应
func DoAddMusic (req *model.AddMusicReq) *model.AddMusicResp {
	resp := &model.AddMusicResp{
		Code: common.HandlerSuccess,
		Msg: "add music success",
		MID: 0,
	}
	music, err := db.AddMusic(req.ToMusic())	// 持久化音乐至DB
	if err != nil {
		resp.Code = common.HandlerDBInsertErr
		resp.Msg = fmt.Sprintf("db insert music err=%v", err)
		return resp
	}
	resp.MID = music.ID
	albumMusic := &model.AlbumMusic{
		Aid: req.AID,
		Mid: music.ID,
	}
	if _, err := db.AddAlbumMusic(albumMusic); err != nil {	// 创建专辑音乐关系
		resp.Code = common.HandlerDBInsertErr
		resp.Msg = fmt.Sprintf("db insert album music connection err=%v", err)
		return resp
	}
	return resp
}

// DoMdfMusic 修改音乐信息
//
// 入参
//	req *model.MdfMusicReq	// 修改音乐请求
// 返回
//	*model.MdfMusicResp		// 修改音乐响应
func DoMdfMusic (req *model.MdfMusicReq) *model.MdfMusicResp {
	resp := &model.MdfMusicResp{
		Code: common.HandlerSuccess,
		Msg: "delete album success",
	}
	if err := db.MdfMusic(req.ToMusic()); err != nil {	// 修改音乐信息
		resp.Code = common.HandlerDBUpdateErr
		resp.Msg = fmt.Sprintf("db update music info err=%v", err)
		return resp
	}
	return resp
}

// DoDeleteMusic 删除音乐
// 只删除专辑音乐关系，不删除音乐
//
// 入参
//	req *model.DeleteMusicReq	// 删除音乐请求
// 返回
//	*model.DeleteMusicResp		// 删除音乐响应
func DoDeleteMusic (req *model.DeleteMusicReq) *model.DeleteMusicResp {
	resp := &model.DeleteMusicResp{
		Code: common.HandlerSuccess,
		Msg: "delete album success",
	}
	if err := db.DelAlbumMusic(req.ToAlbumMusic()); err != nil {	// 删除专辑音乐关系
		resp.Code = common.HandlerDBDeleteErr
		resp.Msg = fmt.Sprintf("db delete album music connection err=%v", err)
		return resp
	}
	return resp
}
