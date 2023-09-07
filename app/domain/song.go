package domain

import "github.com/labstack/echo/v4"

type Song struct {
	ID     uint   `json:"id,omitempty" gorm:"primaryKey"`
	Title  string `json:"title"`
	Album  string `json:"album"`
	Singer string `json:"singer"`
	URL    string `json:"url"`
	Lyrics string `json:"lyrics"`
}

type SongRepository interface {
	Get(ctx echo.Context) (songs []Song, err error)
	GetById(ctx echo.Context, id int64) (song Song, err error)
	Post(ctx echo.Context, song Song) (err error)
	Patch(ctx echo.Context, song Song) (err error)
	Delete(ctx echo.Context, song Song) (err error)
}

type SongUsecase interface {
	Get(ctx echo.Context) (songs []Song, err error)
	GetById(ctx echo.Context, id int64) (song Song, err error)
	Patch(ctx echo.Context, song Song) (err error)
	Delete(ctx echo.Context, song Song) (err error)
	RegisterSong(ctx echo.Context, song Song) (err error)
}
