package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_echoLog "github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"spotipeng/app/global"
	"spotipeng/app/modules/song/http_delivery_song"
	"spotipeng/app/modules/song/song_repository"
	"spotipeng/app/modules/song/song_usecase"
	"spotipeng/app/modules/user/http_delivery_users"
	"spotipeng/app/modules/user/users_repository"
	"spotipeng/app/modules/user/users_usecase"
	"spotipeng/app/util"
)

func init() {
	logrus.Info("Initializing app")

	util.MustReadYaml("/home/rizki/GolangProjects/spotipengdb/app/config/config.yaml")
}

func main() {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", global.Config.Database.User,
		global.Config.Database.Password, global.Config.Database.Host, global.Config.Database.Port, global.Config.Database.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	logLevel := logger.Info

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second * 2,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		},
	)
	dbConn, err := gorm.Open(mysql.Open(connection), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	global.DbConn = dbConn

	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Host,
		Password: global.Config.Redis.Password,
		DB:       0,
	})
	global.RedisClient = rdb

	global.Validate = validator.New(validator.WithRequiredStructEnabled())

	global.Echo = echo.New()

	if true {
		global.Echo.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			util.LoggerI(c, "DEBUG << ", string(reqBody))
			util.LoggerI(c, "DEBUG >> ", string(resBody))
		}))
	}

	global.Echo.Use(middleware.CORS())
	global.Echo.HideBanner = true
	global.Echo.Logger.SetLevel(_echoLog.INFO)
	//global.Echo.Use(echojwt.WithConfig(echojwt.Config{
	//	SigningKey: []byte(global.Config.Encryption.JwtSecret),
	//}))
	global.Echo.HideBanner = true

	registerRepo()

	registerUsecase()

	registerHTTPHandler()

	addr := fmt.Sprintf("%s:%d", global.Config.WebServer.Bind, 8080)
	logrus.Info("Starting HTTP server at ", addr)
	logrus.Fatal(global.Echo.Start(fmt.Sprintf(addr)))

}

func registerRepo() {
	global.UserRepo = users_repository.New()
	global.SongRepo = song_repository.New()
}

func registerUsecase() {
	global.UserUsecase = users_usecase.New()
	global.SongUsecase = song_usecase.New()
}

func registerHTTPHandler() {
	http_delivery_users.HttpUserHandler()
	http_delivery_song.HttpSongHandler()
}
