package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"spotipeng/handler"
	"spotipeng/model"
	"spotipeng/utils"
)

func main() {
	// Replace with your actual MariaDB connection details
	dsn := "root:1234@tcp(localhost:3306)/spotipengdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate will create tables for your models if they don't exist
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Song{})

	// Initialize Echo
	e := echo.New()

	// Apply middleware to group
	protectedGroup := e.Group("")
	protectedGroup.Use(utils.Authenticate)

	// Set up routes
	e.POST("/register", handler.RegisterHandler(db))
	e.POST("/login", handler.LoginHandler(db))
	protectedGroup.GET("/users", handler.GetAllUsersHandler(db))
	protectedGroup.GET("/user", handler.GetUserByIDHandler(db))
	protectedGroup.DELETE("/user/:id", handler.DeleteUserByIDHandler(db))
	protectedGroup.POST("/addsong", handler.AddSongHandler(db))
	protectedGroup.GET("/song", handler.GetSongHandler(db))
	// Start the server
	e.Start(":8080")
}
