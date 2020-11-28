package model

type Music struct {
	ID int	`gorm:"column:id" json:"id"`
	Name string	`gorm:"column:name" json:"name"`
	Poster string `gorm:"column:poster" json:"poster"`
	Path string `gorm:"column:path" json:"path"`
	Author string `gorm:"column:author" json:"author"`
}

type AllMusicResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	MusicList []*Music `json:"music_list"`
}

type AddMusicReq struct {
	AID int `json:"aid"`
	Name string `json:"name"`
	Poster string `json:"poster"`
	Path string `json:"path"`
	Author string `json:"author"`
}

func (req *AddMusicReq) ToMusic () *Music {
	return &Music{
		Name: req.Name,
		Poster: req.Poster,
		Path: req.Path,
		Author: req.Author,
	}
}

type AddMusicResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	MID int `json:"mid"`
}

type MdfMusicReq struct {
	MID int `json:"mid"`
	Name string `json:"name"`
	Poster string `json:"poster"`
	Path string `json:"path"`
	Author string `json:"author"`
}

func (req *MdfMusicReq) ToMusic () *Music {
	return &Music{
		Name: req.Name,
		Poster: req.Poster,
		Path: req.Path,
		Author: req.Author,
	}
}

type MdfMusicResp struct {
	Code int `json:"code"`
	Msg string `json:"'msg'"`
}

type DeleteMusicReq struct {
	AID int `json:"aid"`
	MID int `json:"mid"`
}

func (req *DeleteMusicReq) ToAlbumMusic () *AlbumMusic {
	return &AlbumMusic{
		Aid: req.AID,
		Mid: req.MID,
	}
}

type DeleteMusicResp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}
