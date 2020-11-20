package db

import "fmt"

// AddAlbumMusic 插入歌单歌曲关系
//
// 入参
//	albumMusic *AlbumMusic	// 目标歌单歌曲关系信息
// 返回
//	error		// 错误信息
func AddAlbumMusic(albumMusic *AlbumMusic) error {
	if albumMusic.Mid == 0 || albumMusic.Aid == 0 {	// 判断歌单歌曲关系信息是否完整
		return fmt.Errorf("AlbumMusic:missing require parameters")
	}
	DB.Create(albumMusic)	// 执行插入操作
	if _, err := GetAlbumMusic(albumMusic); err != nil {	// 查询插入后目标歌单歌曲关系是否存在
		return err
	}
	return nil
}

// GetAlbumMusic 查询歌单歌曲关系信息
//
// 入参
//	albumMusic *AlbumMusic	// 目标歌单歌曲关系信息
// 返回
//	*AlbumMusic		// 目标歌单歌曲关系完整信息
//	error		// 错误信息
func GetAlbumMusic(albumMusic *AlbumMusic) (*AlbumMusic, error){
	var selectResp []AlbumMusic
	if albumMusic.Mid == 0 || albumMusic.Aid == 0 {	// 判断歌单歌曲关系信息是否完整
		return nil, fmt.Errorf("AlbumMusic:missing require parameters")
	}
	if albumMusic.ID != 0 {	// 根据ID查询
		DB.Table("user_album").Where("id=?", albumMusic.ID).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query id=%d, resp no datas", albumMusic.ID)
		}
	} else if albumMusic.Mid != 0  { //	根据Mid查询
		DB.Table("user_album").Where("mid=?", albumMusic.Mid).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query Mid=%d, resp no datas", albumMusic.Mid)
		}
	} else {	// 根据aid查询
		DB.Table("user_album").Where("aid=?", albumMusic.Aid).Select([]string{"id", "mid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("AlbumMusic:query aid=%d, resp no datas", albumMusic.Aid)
		}
	}
	return &selectResp[0], nil
}

// MdfAlbumMusic 修改歌单歌曲关系信息
//
// 入参
//	AlbumMusic *AlbumMusic	// 目标歌单歌曲关系信息
// 返回
//	error		// 错误信息
func MdfAlbumMusic(albumMusic *AlbumMusic) error {
	var selectResp []AlbumMusic
	if albumMusic.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("AlbumMusic:missing require parameters")
	}
	if albumMusic.Mid != 0 {	// 更新Mid
		if err := DB.Model(&selectResp).Where("id=?", albumMusic.ID).Update("mid", albumMusic.Mid).Error; err != nil {
			return err
		}
	}
	if albumMusic.Aid != 0 {	// 更新Aid
		if err := DB.Model(&selectResp).Where("id=?", albumMusic.ID).Update("aid", albumMusic.Aid).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelAlbumMusic 删除歌单歌曲关系
//
// 入参
//	AlbumMusic *AlbumMusic	// 目标歌单歌曲关系信息
// 返回
//	error		// 错误信息
func DelAlbumMusic(albumMusic *AlbumMusic) error {
	var selectResp []AlbumMusic
	if albumMusic.ID != 0 {	// 根据ID删除
		if err := DB.Model(&selectResp).Where("id=?", albumMusic.ID).Delete(&AlbumMusic{}).Error; err != nil {
			return err
		}
	} else if albumMusic.Mid != 0 {	// 根据Mid删除
		if err := DB.Model(&selectResp).Where("mid=?", albumMusic.Mid).Delete(&AlbumMusic{}).Error; err != nil {
			return err
		}
	} else {	// 根据aid删除
		if err := DB.Model(&selectResp).Where("aid=?", albumMusic.Aid).Delete(&AlbumMusic{}).Error; err != nil {
			return err
		}
	}
	return nil
}
