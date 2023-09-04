package handler

// import (
// 	"net/http"
// 

// 	"github.com/labstack/echo/v4"
// 	"gorm.io/gorm"

// 	"spotipeng/model"
// )

// func AddSongHandler(db *gorm.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		banner := new(model.Song)
// 		if err := c.Bind(banner); err != nil {
// 			return c.JSON(http.StatusBadRequest, "Invalid request")
// 		}

// 		// Create the new song
// 		newSong := model.Song{
// 			Title:  banner.Title,
// 			Album:  banner.Album,
// 			Singer: banner.Singer,
// 			URL:    banner.URL,
// 		}
// 		db.Create(&newSong)

// 		return c.JSON(http.StatusCreated, "Song registered successfully")
// 	}
// }

// func GetSongHandler(db *gorm.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		var songs []model.Song
// 		result := db.Find(&songs)
// 		if result.Error != nil {
// 			return c.JSON(http.StatusInternalServerError, "Error fetching songs")
// 		}

// 		return c.JSON(http.StatusOK, songs)
// 	}
// }
