package db

import "fmt"

// AddUserAlbum 插入专辑用户关系
//
// 入参
//	userAlbum *UserAlbum	// 目标专辑用户关系信息
// 返回
//	userAlbum *UserAlbum	// 目标专辑用户关系信息
//	error		// 错误信息
func AddUserAlbum(userAlbum *UserAlbum) (*UserAlbum, error) {
	var selectResp []UserAlbum
	if userAlbum.Uid == 0 || userAlbum.Aid == 0 {	// 判断专辑用户关系信息是否完整
		return nil, fmt.Errorf("UserAlbum:missing require parameters")
	}
	DB.Create(userAlbum).Select([]string{"id", "uid", "aid"}).Find(&selectResp)	// 执行插入操作
	if len(selectResp) == 0 {
		return nil, fmt.Errorf("UserAlbum:instert user_album uid=%d aid=%d fail, resp no datas", userAlbum.Uid, userAlbum.Aid)
	}
	return &selectResp[0], nil
}

// GetUserAlbum 查询专辑用户关系信息
//
// 入参
//	userAlbum *UserAlbum	// 目标专辑用户关系信息
// 返回
//	*UserAlbum		// 目标专辑用户关系完整信息
//	error		// 错误信息
func GetUserAlbum(userAlbum *UserAlbum) (*UserAlbum, error){
	var selectResp []UserAlbum
	if userAlbum.Uid == 0 && userAlbum.Aid == 0 {	// 判断专辑用户关系信息是否完整
		return nil, fmt.Errorf("UserAlbum:missing require parameters")
	}
	if userAlbum.ID != 0 {	// 根据ID查询
		DB.Table("user_album").Where("id=?", userAlbum.ID).Select([]string{"id", "uid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("UserAlbum:query id=%d, resp no datas", userAlbum.ID)
		}
	} else if userAlbum.Uid != 0  { //	根据uid查询
		DB.Table("user_album").Where("uid=?", userAlbum.Uid).Select([]string{"id", "uid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("UserAlbum:query uid=%d, resp no datas", userAlbum.Uid)
		}
	} else {	// 根据aid查询
		DB.Table("user_album").Where("aid=?", userAlbum.Aid).Select([]string{"id", "uid", "aid"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("UserAlbum:query aid=%d, resp no datas", userAlbum.Aid)
		}
	}
	return &selectResp[0], nil
}

// MdfUserAlbum 修改专辑用户关系信息
//
// 入参
//	UserAlbum *UserAlbum	// 目标专辑用户关系信息
// 返回
//	error		// 错误信息
func MdfUserAlbum(userAlbum *UserAlbum) error {
	if userAlbum.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("UserAlbum:missing require parameters")
	}
	if userAlbum.Uid != 0 {	// 更新Uid
		if err := DB.Model(userAlbum).Update("uid", userAlbum.Uid).Error; err != nil {
			return err
		}
	}
	if userAlbum.Aid != 0 {	// 更新Aid
		if err := DB.Model(userAlbum).Update("aid", userAlbum.Aid).Error; err != nil {
			return err
		}
	}
	return nil
}

// DelUserAlbum 删除专辑用户关系
//
// 入参
//	UserAlbum *UserAlbum	// 目标专辑用户关系信息
// 返回
//	error		// 错误信息
func DelUserAlbum(userAlbum *UserAlbum) error {
	if userAlbum.ID != 0 {	// 根据ID删除
		if err := DB.Where("id=?", userAlbum.ID).Delete(&UserAlbum{}).Error; err != nil {
			return err
		}
	} else if userAlbum.Uid != 0 {	// 根据uid删除
		if err := DB.Where("uid=?", userAlbum.Uid).Delete(&UserAlbum{}).Error; err != nil {
			return err
		}
	} else {	// 根据aid删除
		if err := DB.Where("aid=?", userAlbum.Aid).Delete(&UserAlbum{}).Error; err != nil {
			return err
		}
	}
	return nil
}
