package song_repository

import (
	"github.com/labstack/echo/v4"

	"spotipeng/app/domain"
	"spotipeng/app/global"
)

type gormSongRepo struct {
}

func (g gormSongRepo) Get(ctx echo.Context) (users []domain.Song, err error) {
	err = global.DbConn.Omit("password").Find(&users).Error
	return
}

func (g gormSongRepo) GetById(ctx echo.Context, id int64) (user domain.Song, err error) {
	err = global.DbConn.Where("id=?", id).First(&user).Error
	return
}

func (g gormSongRepo) GetByEmail(ctx echo.Context, email string) (user domain.Song, err error) {
	err = global.DbConn.Where("email=?", email).First(&user).Error
	return
}

func (g gormSongRepo) Post(ctx echo.Context, user domain.Song) (err error) {
	err = global.DbConn.Create(&user).Error
	return
}

func (g gormSongRepo) Patch(ctx echo.Context, user domain.Song) (err error) {
	err = global.DbConn.Updates(&user).Error
	return
}

func (g gormSongRepo) Delete(ctx echo.Context, user domain.Song) (err error) {
	err = global.DbConn.Model(domain.Song{}).Where("id", user.ID).Update("status", -1).Error
	return
}

func New() domain.SongRepository {
	return &gormSongRepo{}
}
