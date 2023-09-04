package global

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"spotipeng/app/domain"
)

var (
	DbConn      *gorm.DB
	RedisClient *redis.Client
	Echo        *echo.Echo
	Config      domain.Config
	Validate    *validator.Validate
)

var (
	UserRepo domain.UserRepository
	SongRepo domain.SongRepository
)

var (
	UserUsecase domain.UserUsecase
	SongUsecase domain.SongUsecase
)
