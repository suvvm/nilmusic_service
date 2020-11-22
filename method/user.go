package method

import (
	"fmt"
	"suvvm.work/nilmusic_service/common"
	"suvvm.work/nilmusic_service/dal/db"
	"suvvm.work/nilmusic_service/model"
)

// DoLogin 登录业务逻辑方法
// 根据请求参数在DB中查询用户信息，验证密码正确性，构造响应参数
//
// 入参
//	req *model.LoginReq	// 登录请求
// 返回
//	*model.LoginResp	// 登录响应
func DoLogin(req *model.LoginReq) *model.LoginResp {
	resp := &model.LoginResp{}
	user, err := db.GetUser(req.ToUser())
	if err != nil {
		resp.Code = common.HandlerDBSelectErr
		resp.Msg = fmt.Sprintf("%v", err)
		resp.UID = 0
		return resp
	}
	if user.Password != req.Password {	// 验证密码
		resp.Code = common.HandlerPasswordErr
		resp.Msg = "phone number or password error"
		resp.UID = 0
		return resp
	}
	resp.Code = common.HandlerSuccess
	resp.Msg = "login success"
	resp.UID = user.ID
	return resp
}

// DoRegister 注册业务逻辑
// 根据请求参数，将用户信息持久化至DB，构造响应参数
//
// 入参
//	req *model.RegisterReq	// 注册请求
// 返回
//	*model.RegisterResp		// 注册响应
func DoRegister(req *model.RegisterReq) *model.RegisterResp {
	resp := &model.RegisterResp{}
	_, err := db.AddUser(req.ToUser())	// 注册用户至DB
	if err != nil {
		resp.Code = common.HandlerDBInsertErr
		resp.Msg = fmt.Sprintf("%v", err)
		return resp
	}
	resp.Code = common.HandlerSuccess
	resp.Msg = "register success"
	return resp
}
