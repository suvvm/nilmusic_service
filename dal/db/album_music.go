package db

import (
	"fmt"
	"suvvm.work/nilmusic_service/model"
)

// AddAlbumMusic 插入专辑歌曲关系
//
// 入参
//	albumMusic *AlbumMusic	// 目标专辑歌曲关系信息
// 返回
//	albumMusic *AlbumMusic	// 目标专辑歌曲关系信息
//	error		// 错误信息
func AddAlbumMusic(albumMusic *model.AlbumMusic) (*model.AlbumMusic, error) {
	if albumMusic.Mid == 0 || albumMusic.Aid == 0 {	// 判断专辑歌曲关系信息是否完整
		return nil, fmt.Errorf("AlbumMusic:missing require parameters")
	}
	DB.Create(albumMusic)	// 执行插入操作
	return albumMusic, nil
}

// GetAlbumMusic 查询专辑歌曲关系信息
//
// 入参
//	albumMusic *AlbumMusic	// 目标专辑歌曲关系信息
// 返回
//	*AlbumMusic		// 目标专辑歌曲关系完整信息
//	error		// 错误信息
func GetAlbumMusic(albumMusic *model.AlbumMusic) (*[]model.AlbumMusic, error){
	var selectResp []model.AlbumMusic
	if albumMusic.Mid == 0 && albumMusic.Aid == 0 {	// 判断专辑歌曲关系信息是否完整
		return nil, fmt.Errorf("AlbumMusic:missing require parameters")
	}
	if albumMusic.ID != 0 {	// 根据ID查询
		DB.Table("album_music").Where("id=?", albumMusic.ID).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query id=%d, resp no datas", albumMusic.ID)
		}
	} else if albumMusic.Mid != 0  { //	根据Mid查询
		DB.Table("album_music").Where("mid=?", albumMusic.Mid).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query Mid=%d, resp no datas", albumMusic.Mid)
		}
	} else {	// 根据aid查询
		DB.Table("album_music").Where("aid=?", albumMusic.Aid).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query aid=%d, resp no datas", albumMusic.Aid)
		}
	}
	return &selectResp, nil
}

// MdfAlbumMusic 修改专辑歌曲关系信息
//
// 入参
//	AlbumMusic *AlbumMusic	// 目标专辑歌曲关系信息
// 返回
//	error		// 错误信息
func MdfAlbumMusic(albumMusic *model.AlbumMusic) error {
	if albumMusic.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("AlbumMusic:missing require parameters")
	}
	if albumMusic.Mid != 0 {	// 更新Mid
		if err := DB.Model(albumMusic).Update("mid", albumMusic.Mid).Error; err != nil {
			return err
		}
	}
	if albumMusic.Aid != 0 {	// 更新Aid
		if err := DB.Model(albumMusic).Update("aid", albumMusic.Aid).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelAlbumMusic 删除专辑歌曲关系
//
// 入参
//	AlbumMusic *AlbumMusic	// 目标专辑歌曲关系信息
// 返回
//	error		// 错误信息
func DelAlbumMusic(albumMusic *model.AlbumMusic) error {
	if albumMusic.Mid != 0 && albumMusic.Aid != 0 {
		if err := DB.Where("mid=? AND aid=?", albumMusic.Mid, albumMusic.Aid).Delete(&model.AlbumMusic{}).Error; err != nil {
			return err
		}
	}else if albumMusic.ID != 0 {	// 根据ID删除
		if err := DB.Where("id=?", albumMusic.ID).Delete(&model.AlbumMusic{}).Error; err != nil {
			return err
		}
	} else if albumMusic.Mid != 0 {	// 根据Mid删除
		if err := DB.Where("mid=?", albumMusic.Mid).Delete(&model.AlbumMusic{}).Error; err != nil {
			return err
		}
	} else {	// 根据aid删除
		if err := DB.Where("aid=?", albumMusic.Aid).Delete(&model.AlbumMusic{}).Error; err != nil {
			return err
		}
	}
	return nil
}
