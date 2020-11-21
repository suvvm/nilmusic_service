package db

import "fmt"

// AddAlbum 插入专辑
//
// 入参
//	album *Album	// 目标专辑信息
// 返回
//	album *Album	// 目标专辑信息
//	error			// 错误信息
func AddAlbum(album *Album) (*Album, error) {
	var selectResp []Album
	if album.Name == ""{	// 判断专辑信息是否完整
		return nil, fmt.Errorf("album:missing require parameters")
	}
	DB.Create(album).Select([]string{"id", "name", "poster", "playnum"}).Find(&selectResp)	// 执行插入操作
	if len(selectResp) == 0 {
		return nil, fmt.Errorf("album:instert user name=%s fail, resp no datas", album.Name)
	}
	return &selectResp[0], nil
}

// GetAlbum 查询专辑信息
//
// 入参
//	album *Album	// 目标专辑信息
// 返回
//	*[]Album		// 查询结果
//	error			// 错误信息
func GetAlbum(album *Album) (*[]Album, error){
	var selectResp []Album
	if album.ID != 0 {	// 根据ID查询
		DB.Table("albums").Where("id=?", album.ID).Select([]string{"id", "name", "poster", " playnum"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("album:query id=%d, resp no datas", album.ID)
		}
	} else if album.Name != "" { //	根据name查询
		DB.Table("albums").Where("name=?", album.Name).Select([]string{"id", "name", "poster", " playnum"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("album:query name=%s, resp no datas", album.Name)
		}
	} else {	// 根据封面查询
		DB.Table("albums").Where("poster=?", album.Poster).Select([]string{"id", "name", "poster", " playnum"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("album:query poster=%s, resp no datas", album.Poster)
		}
	}
	return &selectResp, nil
}

// MdfAlbum 修改专辑信息
//
// 入参
//	album *Album	// 目标专辑信息
// 返回
//	error			// 错误信息
func MdfAlbum(album *Album) error {
	if album.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("album:missing require parameters")
	}
	if album.Name != "" {	// 更新专辑名
		if err := DB.Model(album).Update("name", album.Name).Error; err != nil {
			return err
		}
	}
	if album.Poster != "" {	// 更新专辑封面
		if err := DB.Model(album).Update("poster", album.Poster).Error; err != nil {
			return err
		}
	}
	if album.Playnum != "" {	// 更新专辑封面
		if err := DB.Model(album).Update("playnum", album.Playnum).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelAlbum 删除专辑
//
// 入参
//	album *Album	// 目标专辑信息
// 返回
//	error			// 错误信息
func DelAlbum(album *Album) error {
	if album.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("album:missing require parameters")
	}
	if err := DB.Where("id=?", album.ID).Delete(&Album{}).Error; err != nil {
		return err
	}
	return nil
}


