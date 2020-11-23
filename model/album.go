package model

type Album struct {
	ID int	`gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Poster string `gorm:"column:poster"`
	Playnum string `gorm:"column:playnum"`
}

type AlbumMusic struct {
	ID int	`gorm:"column:id"`
	Mid int `gorm:"column:mid"`
	Aid int `gorm:"column:aid"`
}

func (AlbumMusic) TableName() string {
	return "album_music"
}

type AllAlbumResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	AlbumList []*Album `json:"album_list"`
}

type CreateAlbumReq struct {
	UID int `json:"uid"`
	Name string `json:"name"`
	Poster string `json:"poster"`
	Playnum string `json:"playnum"`
}

func (req *CreateAlbumReq) ToAlbum() *Album {
	return &Album{
		Name: req.Name,
		Poster: req.Poster,
		Playnum: req.Playnum,
	}
}

type CreateAlbumResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}

type DeleteAlbumReq struct {
	UID int `json:"uid"`
	AID int `json:"aid"`
}

func (req *DeleteAlbumReq) ToUserAlbum () *UserAlbum {
	return &UserAlbum{
		Uid: req.UID,
		Aid: req.AID,
	}
}

type DeleteAlbumResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
