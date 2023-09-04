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

	// Read the YAML file
	util.MustReadYaml("/home/rizki/GolangProjects/spotipengdb/app/config/spotipeng/config.yaml")

}

func main() {

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", global.Config.Database.User,
		global.Config.Database.Password, global.Config.Database.Host, global.Config.Database.Port, global.Config.Database.Name)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	//DATABASE
	logLevel := logger.Info

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  logLevel,        // Log level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,            // Disable color
		},
	)
	dbConn, err := gorm.Open(mysql.Open(connection), &gorm.Config{Logger: newLogger})
	if err != nil {
		logrus.Error(err.Error())
		panic(err)
	}
	global.DbConn = dbConn

	//REDIS
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Host,
		Password: global.Config.Redis.Password, // no password set
		DB:       0,                            // use default DB
	})
	global.RedisClient = rdb

	//VALIDATOR
	//we make it single instance, because it does cache the struct schema for better performance
	global.Validate = validator.New(validator.WithRequiredStructEnabled())

	//ECHO
	global.Echo = echo.New()

	//Body Dump for debugging purpose, it will show request response body
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

	//Register Repo
	registerRepo()

	//Register Usecase
	registerUsecase()

	//Register Handler
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
	// Register HTTP Handler here, when inserting new handler, please order by alphabetical for easier reading
	http_delivery_users.HttpUserHandler()
	http_delivery_song.HttpSongHandler()
}

// var (
// 	qCreateSchema string
// )

// func main() {
// 	//To run: go run . -up
// 	util.MustReadYaml(util.GetEnvSetIfEmpty("CONFIG_FILE", "/app/config/config.yaml"))
// 	createSchema(global.Config.Database.RootPassword, global.Config.Database.Host)
// 	migrate(global.Config.Database.User, global.Config.Database.Password, global.Config.Database.Host)
// }

// func createSchema(rootPass, host string) {
// 	db := sqlx.MustConnect(mysqlDbConfig(rootPass, host))
// 	_, err := db.Exec(qCreateSchema)
// 	iferr.Fatal(err)
// }

// func migrate(dbUser, dbPass, host string) {
// 	db := sqlx.MustConnect(stellarDbConfig(dbUser, dbPass, host))
// 	vmigrate.Migrate(db, migrationHandlers)
// }

// import (
// 	"github.com/labstack/echo/v4"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"

// 	"spotipeng/app/domain"
// 	"spotipeng/app/handler"
// 	"spotipeng/app/util"
// )

// func main() {
// 	// Replace with your actual MariaDB connection details
// 	dsn := "root:1234@tcp(localhost:3306)/spotipengdb?charset=utf8mb4&parseTime=True&loc=Local"

// 	// Connect to the database
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("Failed to connect to the database")
// 	}

// 	// AutoMigrate will create tables for your models if they don't exist
// 	db.AutoMigrate(&domain.User{})
// 	db.AutoMigrate(&domain.Song{})

// 	// Initialize Echo
// 	e := echo.New()

// 	// Apply middleware to group
// 	protectedGroup := e.Group("")
// 	protectedGroup.Use(util.Authenticate)

// 	// Set up routes
// 	e.POST("/register", handler.RegisterHandler(db))
// 	e.POST("/login", handler.LoginHandler(db))
// 	protectedGroup.GET("/users", handler.GetAllUsersHandler(db))
// 	protectedGroup.GET("/user", handler.GetUserByIDHandler(db))
// 	protectedGroup.DELETE("/user/:id", handler.DeleteUserByIDHandler(db))
// 	protectedGroup.POST("/addsong", handler.AddSongHandler(db))
// 	protectedGroup.GET("/song", handler.GetSongHandler(db))
// 	// Start the server
// 	e.Start(":8080")
// }
