package method

import (
	"fmt"
	"log"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/dal/db"
	"suvvm.work/nilmusic_service/model"
)

// DoGetAllAlbum 获取用户全部专辑业务逻辑
// 判断用户uid是否获取成功，并根据uid获取对应专辑列表，并构造获取专辑响应
//
// 入参
//	uid int				// 用户uid
// 返回
//	*model.AllAlbumResp	// 获取专辑响应
func DoGetAllAlbum (uid int) *model.AllAlbumResp {
	resp := &model.AllAlbumResp{
		Code: common.HandlerSuccess,
		Msg: "get album success",
		AlbumList: make([]*model.Album, 0),
	}
	if uid == 0 {
		resp.Code = common.HandlerReadPathErr
		resp.Msg = common.HandlerReadPathErrMsg
		return resp
	}
	albumList, err := GetAlbumByUID(uid)	// 根据uid获取专辑列表
	if err != nil {
		resp.Code = common.HandlerDBSelectErr
		resp.Msg = fmt.Sprintf("db get album list err=%v", err)
		return resp
	}
	resp.AlbumList = albumList
	return resp
}

// GetAlbumByUID 根据用户ID获取目标用户创建的全部专辑
// 根据用户ID查询到专辑AID列表后，遍历AID获取专辑信息并返回专辑列表
//
// 入参
//	uid int			// 用户uid
// 返回
//	[]*model.Album	// 专辑列表
//	error			// 错误信息
func GetAlbumByUID (uid int)  ([]*model.Album, error) {
	albumList := make([]*model.Album, 0)
	userAlbum := &model.UserAlbum{Uid: uid}
	userAlbumList, err := db.GetUserAlbum(userAlbum)	// user_album获取uid对应的全部关系
	if err != nil {
		return nil, err
	}
	for _, userAlbum := range *userAlbumList {	// 遍历关系，根据aid获取专辑详细信息
		album := &model.Album{ID: userAlbum.Aid}
		albums, err := db.GetAlbum(album)
		if err != nil {
			log.Printf("GetAlbumByUID: get album by AID=%d err=%v", userAlbum.Aid, err)
			continue
		}
		for index := range *albums {	// 追加专辑至albumList
			albumList = append(albumList, &(*albums)[index])
		}
	}
	return albumList, nil
}

// DoCreateAlbum 创建专辑
// 创建专辑并添加目标用户与新建专辑的关系
//
// 入参
//	req *model.CreateAlbumReq	// 创建专辑请求
// 返回
//	*model.CreateAlbumResp		// 创建专辑响应
func DoCreateAlbum (req *model.CreateAlbumReq) *model.CreateAlbumResp {
	resp := &model.CreateAlbumResp{
		Code: common.HandlerSuccess,
		Msg: "create album success",
		AID: 0,
	}
	album, err := db.AddAlbum(req.ToAlbum())	// 持久化专辑至DB
	if err != nil {
		resp.Code = common.HandlerDBInsertErr
		resp.Msg = fmt.Sprintf("db insert album err=%v", err)
		return resp
	}
	resp.AID = album.ID
	userAlbum := &model.UserAlbum{
		Uid: req.UID,
		Aid: album.ID,
	}
	if _, err := db.AddUserAlbum(userAlbum); err != nil {	// 创建用户专辑关系
		resp.Code = common.HandlerDBInsertErr
		resp.Msg = fmt.Sprintf("db insert user album connection err=%v", err)
		return resp
	}
	return resp
}

// DoDeleteAlbum 删除专辑
// 只删除用户专辑关系，不删除专辑
//
// 入参
//	req *model.DeleteAlbumReq	// 删除专辑请求
// 返回
//	*model.DeleteAlbumResp		// 删除专辑响应
func DoDeleteAlbum (req *model.DeleteAlbumReq) *model.DeleteAlbumResp {
	resp := &model.DeleteAlbumResp{
		Code: common.HandlerSuccess,
		Msg: "delete album success",
	}
	if err := db.DelUserAlbum(req.ToUserAlbum()); err != nil {	// 删除用户专辑关系
		resp.Code = common.HandlerDBDeleteErr
		resp.Msg = fmt.Sprintf("db delete user album connection err=%v", err)
		return resp
	}
	return resp
}