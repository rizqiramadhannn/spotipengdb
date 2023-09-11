package users_repository

import (
	"github.com/labstack/echo/v4"

	"spotipeng/app/domain"
	"spotipeng/app/global"
)

type gormUserRepo struct {
}

func (g gormUserRepo) Get(ctx echo.Context) (users []domain.User, err error) {
	err = global.DbConn.Omit("password").Find(&users).Error
	return
}

func (g gormUserRepo) GetById(ctx echo.Context, id int64) (user domain.User, err error) {
	err = global.DbConn.Where("id=?", id).First(&user).Error
	return
}

func (g gormUserRepo) GetByEmail(ctx echo.Context, email string) (user domain.User, err error) {
	err = global.DbConn.Where("email=?", email).First(&user).Error
	return
}

func (g gormUserRepo) Post(ctx echo.Context, user domain.User) (err error) {
	err = global.DbConn.Create(&user).Error
	return
}

func (g gormUserRepo) Patch(ctx echo.Context, user domain.User) (err error) {
	err = global.DbConn.Updates(&user).Error
	return
}

func (g gormUserRepo) Delete(ctx echo.Context, user domain.User) (err error) {
	err = global.DbConn.Delete(&user).Error
	return
}

func New() domain.UserRepository {
	return &gormUserRepo{}
}
