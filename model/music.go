package model

type Music struct {
	ID int	`gorm:"column:id"`
	Name string	`gorm:"column:name"`
	Poster string `gorm:"column:poster"`
	Path string `gorm:"column:path"`
	Author string `gorm:"column:author"`
}
