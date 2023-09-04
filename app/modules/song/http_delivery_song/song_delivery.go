package http_delivery_song

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"spotipeng/app/domain"
	"spotipeng/app/global"
	"spotipeng/app/modules/http/http_usecase"
	"spotipeng/app/util"
)

type SongHandler struct {
	SongUsecase domain.SongUsecase
}

func (h SongHandler) CreateSong(c echo.Context) error {
	data := new(domain.Song)
	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	}

	if err := global.Validate.Struct(data); err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": err.Error(),
		})
	}

	if err := global.SongUsecase.RegisterSong(c, *data); err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": "Failed to register song",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"rc":  domain.RC_00_OK,
		"msg": "Song registered successfully",
	})
}

func (h SongHandler) ListSong(c echo.Context) error {
	users, err := global.SongUsecase.Get(c)
	if err != nil {
		logrus.Error(err.Error())
		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"rc":    domain.RC_00_OK,
			"users": users,
		})
	}
}

func (h SongHandler) GetSongByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": err.Error(),
		})
	}

	user, err := global.SongUsecase.GetById(c, int64(id))
	if err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"rc":   domain.RC_00_OK,
			"user": user,
		})
	}
}

func HttpSongHandler() {
	handler := &SongHandler{}

	v1 := "/spotipeng/api/v1"
	global.Echo.POST(v1+"/song", handler.CreateSong, http_usecase.IsLoggedIn)
	global.Echo.GET(v1+"/song", handler.ListSong, http_usecase.IsLoggedIn)
	global.Echo.GET(v1+"/song/:id", handler.GetSongByID, http_usecase.IsLoggedIn)
	// global.Echo.POST(v1+"/users/google_login", handler.GoogleLogin)
}
