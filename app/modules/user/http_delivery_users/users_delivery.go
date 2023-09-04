package http_delivery_users

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

type UserHandler struct {
	UserUsecase domain.UserUsecase
}

func (h UserHandler) Login(c echo.Context) error {
	data := new(domain.User)
	if err := c.Bind(data); err != nil {
		return err
	}

	accessToken, refreshToken, err := global.UserUsecase.Login(c, data.Email, data.Password)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"rc":  domain.RC_02_INVALID_AUTHORIZATION,
			"msg": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, map[string]string{
			"rc":            domain.RC_00_OK,
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}

// func (h UserHandler) GoogleLogin(c echo.Context) error {
// 	data := new(domain.GoogleLogin)
// 	if err := c.Bind(data); err != nil {
// 		return c.JSON(http.StatusBadRequest, domain.ErrBadParamInput.Error())
// 	}

// 	accessToken, refreshToken, err := global.UserUsecase.GoogleLogin(c, *data)
// 	if err != nil {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"rc":  domain.RC_02_INVALID_AUTHORIZATION,
// 			"msg": err.Error(),
// 		})
// 	} else {
// 		return c.JSON(http.StatusOK, map[string]string{
// 			"rc":            domain.RC_00_OK,
// 			"access_token":  accessToken,
// 			"refresh_token": refreshToken,
// 		})
// 	}
// }

func (h UserHandler) CreateUser(c echo.Context) error {
	data := new(domain.User)
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

	if err := global.UserUsecase.RegisterUser(c, *data); err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"rc":  domain.RC_03_INTERNAL_ERROR,
			"msg": "Failed to create user",
		})
	}

	// Return a success response when user creation is successful
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"rc":  domain.RC_00_OK,
		"msg": "User created successfully",
	})
}

func (h UserHandler) ListUser(c echo.Context) error {
	users, err := global.UserUsecase.Get(c)
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

func (h UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		util.LoggerI(c, err.Error())
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"rc":  domain.RC_01_INVALID_PAYLOAD,
			"msg": err.Error(),
		})
	}

	user, err := global.UserUsecase.GetById(c, int64(id))
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

func HttpUserHandler() {
	handler := &UserHandler{}

	v1 := "/spotipeng/api/v1"
	global.Echo.POST(v1+"/users/login", handler.Login)
	global.Echo.POST(v1+"/register", handler.CreateUser)
	global.Echo.GET(v1+"/users", handler.ListUser, http_usecase.IsLoggedIn)
	global.Echo.GET(v1+"/users/:id", handler.GetUserByID, http_usecase.IsLoggedIn)
	// global.Echo.POST(v1+"/users/google_login", handler.GoogleLogin)
}
