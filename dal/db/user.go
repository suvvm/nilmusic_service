package db

import "fmt"

// AddUser 插入用户
//
// 入参
//	user *User	// 目标用户信息
// 返回
//	error		// 错误信息
func AddUser(user *User) error {
	if user.Pnum == "" || user.Password == "" {	// 判断用户信息是否完整
		return fmt.Errorf("user:missing require parameters")
	}
	DB.Create(user)	// 执行插入操作
	if _, err := GetUser(user); err != nil {	// 查询插入后目标用户是否存在
		return err
	}
	return nil
}

// GetUser 查询用户信息
//
// 入参
//	user *User	// 目标用户信息
// 返回
//	*User		// 目标用户完整信息
//	error		// 错误信息
func GetUser(user *User) (*User, error){
	var selectResp []User
	if user.ID != 0 {	// 根据ID查询
		DB.Table("user").Where("id=?", user.ID).Select([]string{"id", "pnum"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("user:query id=%d, resp no datas", user.ID)
		}
	} else { //	根据pnum查询
		DB.Table("user").Where("pnum=?", user.Pnum).Select([]string{"id", "pnum"}).Find(&selectResp)
		if len(selectResp) == 0 {
			return nil, fmt.Errorf("user:query pnum=%s, resp no datas", user.Pnum)
		}
	}
	return &selectResp[0], nil
}

// MdfUser 修改用户信息
//
// 入参
//	user *User	// 目标用户信息
// 返回
//	error		// 错误信息
func MdfUser(user *User) error {
	var selectResp []User
	if user.ID == 0 { // 判断ID是否为空
		return fmt.Errorf("user:missing require parameters")
	}
	if user.Pnum != "" {	// 更新手机号
		if err := DB.Model(&selectResp).Where("id=?", user.ID).Update("pnum", user.Pnum).Error; err != nil {
			return err
		}
	}
	if user.Password != "" {	// 更新密码
		if err := DB.Model(&selectResp).Where("id=?", user.ID).Update("password", user.Password).Error; err != nil {
				return err
		}
	}
	return nil
}

// DelUser 删除用户
//
// 入参
//	user *User	// 目标用户信息
// 返回
//	error		// 错误信息
func DelUser(user *User) error {
	var selectResp []User
	if user.ID != 0 {	// 根据ID删除
		if err := DB.Model(&selectResp).Where("id=?", user.ID).Delete(&User{}).Error; err != nil {
			return err
		}
	} else {	// 根据手机号删除
		if err := DB.Model(&selectResp).Where("pnum=?", user.Pnum).Delete(&User{}).Error; err != nil {
			return err
		}
	}
	return nil
}
