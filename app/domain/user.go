package domain

import "github.com/labstack/echo/v4"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type UserRepository interface {
	Get(ctx echo.Context) (users []User, err error)
	GetById(ctx echo.Context, id int64) (user User, err error)
	GetByEmail(ctx echo.Context, email string) (user User, err error)
	Post(ctx echo.Context, user User) (err error)
	Patch(ctx echo.Context, user User) (err error)
	Delete(ctx echo.Context, user User) (err error)
}

type UserUsecase interface {
	Get(ctx echo.Context) (users []User, err error)
	GetById(ctx echo.Context, id int64) (user User, err error)
	Patch(ctx echo.Context, user User) (err error)
	Delete(ctx echo.Context, user User) (err error)
	Login(ctx echo.Context, email, password string) (accessToken string, refreshToken string, err error)
	RegisterUser(ctx echo.Context, user User) (err error)
}
