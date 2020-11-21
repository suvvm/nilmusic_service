package db

type User struct {
	ID int	`gorm:"column:id"`
	Pnum string	`gorm:"column:pnum"`
	Password string `gorm:"column:password"`
}

type Album struct {
	ID int	`gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Poster string `gorm:"column:poster"`
	Playnum string `gorm:"column:playnum"`
}

type Music struct {
	ID int	`gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Poster string `gorm:"column:poster"`
	Path string `gorm:"column:path"`
	Author string `gorm:"column:author"`
}

type UserAlbum struct {
	ID int	`gorm:"column:id"`
	Uid int `gorm:"column:uid"`
	Aid int `gorm:"column:aid"`
}

func (UserAlbum) TableName() string {
	return "user_album"
}

type AlbumMusic struct {
	ID int	`gorm:"column:id"`
	Mid int `gorm:"column:mid"`
	Aid int `gorm:"column:aid"`
}

func (AlbumMusic) TableName() string {
	return "album_music"
}