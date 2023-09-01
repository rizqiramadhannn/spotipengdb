package model

type Song struct {
	ID     uint `gorm:"primaryKey"`
	Title  string
	Album  string
	Singer string
	URL    string
}
