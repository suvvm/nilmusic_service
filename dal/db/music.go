package db

import "fmt"

// AddMusic 插入音乐
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	error			// 错误信息
func AddMusic(music *Music) error {
	if music.Name == "" || music.Path == ""{	// 判断音乐信息是否完整
		return fmt.Errorf("music:missing require parameters")
	}
	DB.Create(music)	// 执行插入操作
	if _, err := GetMusic(music); err != nil {	// 查询插入后目标音乐是否存在
		return err
	}
	return nil
}

// GetMusic 查询音乐信息
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	*[]Music		// 查询结果
//	error			// 错误信息
func GetMusic(music *Music) (*[]Music, error){
	var selectResp []Music
	if music.ID != 0 {	// 根据ID查询
		DB.Table("music").Where("id=?", music.ID).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("music:query id=%d, resp no datas", music.ID)
		}
	} else if music.Name != "" { //	根据name查询
		DB.Table("music").Where("name=?", music.Name).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("music:query name=%s, resp no datas", music.Name)
		}
	} else {	// 根据作者查询
		DB.Table("music").Where("author=?", music.Author).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("music:query author=%s, resp no datas", music.Author)
		}
	}
	return &selectResp, nil
}

// MdfMusic 修改音乐信息
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	error			// 错误信息
func MdfMusic(music *Music) error {
	var selectResp []Music
	if music.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("music:missing require parameters")
	}
	if music.Name != "" {	// 更新音乐名
		if err := DB.Model(&selectResp).Where("id=?", music.ID).Update("name", music.Name).Error; err != nil {
			return err
		}
	}
	if music.Poster != "" {	// 更新音乐封面
		if err := DB.Model(&selectResp).Where("id=?", music.ID).Update("poster", music.Poster).Error; err != nil {
			return err
		}
	}
	if music.Path != "" {	// 更新音乐外链
		if err := DB.Model(&selectResp).Where("id=?", music.ID).Update("path", music.Path).Error; err != nil {
			return err
		}
	}
	if music.Author != "" {	// 更新音乐作者
		if err := DB.Model(&selectResp).Where("id=?", music.ID).Update("author", music.Author).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelMusic 删除音乐
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	error			// 错误信息
func DelMusic(music *Music) error {
	var selectResp []Music
	if music.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("music:missing require parameters")
	}
	if err := DB.Model(&selectResp).Where("id=?", music.ID).Delete(&Music{}).Error; err != nil {
		return err
	}
	return nil
}


