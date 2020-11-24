package db

import (
	"fmt"
	"suvvm.work/nilmusic_service/model"
)

// AddMusic 插入音乐
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	music *Music	// 目标音乐信息
//	error			// 错误信息
func AddMusic(music *model.Music) (*model.Music, error) {
	if music.Name == "" || music.Path == ""{	// 判断音乐信息是否完整
		return nil, fmt.Errorf("music:missing require parameters")
	}
	DB.Create(music)	// 执行插入操作
	return music, nil
}

// GetMusic 查询音乐信息
//
// 入参
//	music *Music	// 目标音乐信息
// 返回
//	*[]Music		// 查询结果
//	error			// 错误信息
func GetMusic(music *model.Music) (*[]model.Music, error){
	var selectResp []model.Music
	if music.ID != 0 {	// 根据ID查询
		DB.Table("musics").Where("id=?", music.ID).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("music:query id=%d, resp no datas", music.ID)
		}
	} else if music.Name != "" { //	根据name查询
		DB.Table("musics").Where("name=?", music.Name).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("music:query name=%s, resp no datas", music.Name)
		}
	} else {	// 根据作者查询
		DB.Table("musics").Where("author=?", music.Author).Select([]string{"id", "name", "poster", "path", "author"}).Find(&selectResp)
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
func MdfMusic(music *model.Music) error {
	if music.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("music:missing require parameters")
	}
	if music.Name != "" {	// 更新音乐名
		if err := DB.Model(music).Update("name", music.Name).Error; err != nil {
			return err
		}
	}
	if music.Poster != "" {	// 更新音乐封面
		if err := DB.Model(music).Update("poster", music.Poster).Error; err != nil {
			return err
		}
	}
	if music.Path != "" {	// 更新音乐外链
		if err := DB.Model(music).Update("path", music.Path).Error; err != nil {
			return err
		}
	}
	if music.Author != "" {	// 更新音乐作者
		if err := DB.Model(music).Update("author", music.Author).Error; err != nil {
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
func DelMusic(music *model.Music) error {
	if music.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("music:missing require parameters")
	}
	if err := DB.Where("id=?", music.ID).Delete(&model.Music{}).Error; err != nil {
		return err
	}
	return nil
}


