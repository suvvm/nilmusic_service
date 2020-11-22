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


