package model

type User struct {
	ID int	`gorm:"column:id" json:"id"`
	Pnum string	`gorm:"column:pnum" json:"pnum"`
	Password string `gorm:"column:password" json:"password"`
}

type UserAlbum struct {
	ID int	`gorm:"column:id"`
	Uid int `gorm:"column:uid"`
	Aid int `gorm:"column:aid"`
}

func (UserAlbum) TableName() string {
	return "user_album"
}

type RegisterReq struct {
	PNum string `json:"pnum"`
	Password string `json:"password"`
}

func (req *RegisterReq) ToUser() *User {
	return &User{
		Pnum: req.PNum,
		Password: req.Password,
	}
}

type RegisterResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type LoginReq struct {
	PNum string `json:"pnum"`
	Password string `json:"password"`
}

func (req *LoginReq) ToUser() *User {
	return &User{
		Pnum: req.PNum,
		Password: req.Password,
	}
}

type LoginResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	UID int `json:"uid"`
}
