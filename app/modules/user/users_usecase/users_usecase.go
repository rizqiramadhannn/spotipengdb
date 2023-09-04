package users_usecase

import (
	"errors"
	"time"

	_jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"spotipeng/app/domain"
	"spotipeng/app/global"
	"spotipeng/app/util"
)

type userUsecase struct {
}

// func (u userUsecase) RegisterReseller(ctx echo.Context, user domain.User) (err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

func (u userUsecase) RegisterUser(ctx echo.Context, user domain.User) (err error) {
	// Check if the email is already registered
	existingUser, err := global.UserRepo.GetByEmail(ctx, user.Email)

	if existingUser != (domain.User{}) {
		// Email is already registered, return an error
		return errors.New("Email already registered")
	}

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// Handle the error (e.g., log it)
		util.LoggerI(ctx, err.Error())
		return err
	}

	// Create a new user object
	newUser := domain.User{
		Email:    user.Email,
		Password: string(hashedPassword),
		Name:     user.Name,
		// Set other user properties as needed
	}

	// Store the new user in the database
	if err := global.UserRepo.Post(ctx, newUser); err != nil {
		// Handle the error (e.g., log it)
		util.LoggerI(ctx, err.Error())
		return err
	}

	// Registration successful
	return nil
}

func (u userUsecase) Get(ctx echo.Context) (users []domain.User, err error) {
	return global.UserRepo.Get(ctx)
}

func (u userUsecase) GetById(ctx echo.Context, id int64) (user domain.User, err error) {
	user, err = global.UserRepo.GetById(ctx, id)
	user.Password = ""
	return
}

func (u userUsecase) Patch(ctx echo.Context, user domain.User) (err error) {
	return global.UserRepo.Patch(ctx, user)
}

func (u userUsecase) Delete(ctx echo.Context, user domain.User) (err error) {
	return global.UserRepo.Delete(ctx, user)
}

func (u userUsecase) Login(ctx echo.Context, email, password string) (accessToken string, refreshToken string, err error) {
	user, err := global.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		util.LoggerI(ctx, err.Error())
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		//Invalid password
		util.LoggerI(ctx, err.Error())
		return
	}

	util.LoggerI(ctx, "Login using email & password success for ", user.Email)

	// Create token
	token := _jwt.New(_jwt.SigningMethodHS256)

	// Set claims access token
	claims := token.Claims.(_jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	// Generate encoded token and send it as response.
	accessToken, err = token.SignedString([]byte("1234"))
	if err != nil {
		util.LoggerI(ctx, err.Error())
		return
	}

	// Set claims refresh token
	claimsRefreshToken := token.Claims.(_jwt.MapClaims)
	claimsRefreshToken["email"] = user.Email
	claimsRefreshToken["id"] = user.ID
	claimsRefreshToken["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	refreshToken, err = token.SignedString([]byte("1234"))
	if err != nil {
		util.LoggerI(ctx, err.Error())
		return
	}
	return
}

// func (u userUsecase) Set2FA(ctx echo.Context, data domain.User, enabled bool) (err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (u userUsecase) GoogleLogin(ctx echo.Context, data domain.GoogleLogin) (accessToken string, refreshToken string, err error) {
// 	payload, err := idtoken.Validate(context.Background(), data.Credential, global.Config.Google.Audience)

// 	if err != nil {
// 		util.LoggerI(ctx, err.Error())
// 		return
// 	}

// 	//Cek is expired
// 	exp := payload.Claims["exp"].(float64)
// 	if time.Now().Unix() >= int64(exp) {
// 		err = errors.New("token expired")
// 		util.LoggerI(ctx, err.Error())
// 		return
// 	}

// 	user, err := global.UserRepo.GetByEmail(ctx, fmt.Sprint(payload.Claims["email"]))
// 	if err != nil {
// 		util.LoggerI(ctx, err.Error())
// 		return
// 	}

// 	if user.Status == 1 {

// 		util.LoggerI(ctx, "Login SUCCESS by "+user.Email+" from IP "+ip.GetIPAddress(ctx)+" "+ip.GetUserAgent(ctx))
// 		// Create token
// 		token := _jwt.New(_jwt.SigningMethodHS256)

// 		//get role
// 		role, errr := global.RoleUsecase.GetById(ctx, user.RolesID)
// 		if errr != nil {
// 			util.LoggerI(ctx, errr.Error())
// 			err = errr
// 			return
// 		}

// 		// Set claims access token
// 		claims := token.Claims.(_jwt.MapClaims)
// 		claims["email"] = user.Email
// 		claims["name"] = user.Name
// 		claims["id"] = user.Id
// 		claims["role"] = role.Name
// 		claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

// 		// Generate encoded token and send it as response.
// 		accessToken, err = token.SignedString([]byte(global.Config.Encryption.JwtSecret))
// 		if err != nil {
// 			util.LoggerI(ctx, err.Error())
// 			return
// 		}

// 		// Set claims refresh token
// 		claimsRefreshToken := token.Claims.(_jwt.MapClaims)
// 		claimsRefreshToken["email"] = user.Email
// 		claimsRefreshToken["id"] = user.Id
// 		claimsRefreshToken["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

// 		refreshToken, err = token.SignedString([]byte(global.Config.Encryption.RefreshTokenSecret))
// 		if err != nil {
// 			util.LoggerI(ctx, err.Error())
// 			return
// 		}
// 		return
// 	}
// 	return
// }

// func (u userUsecase) ChangePassword(ctx echo.Context, user domain.User, oldPassword, newPassword string) (err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (u userUsecase) ResetPassword(ctx echo.Context, email, resetCode, password string) (err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (u userUsecase) VerifyEmailAddress(ctx echo.Context, activationCode string) (token string, err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

// func (u userUsecase) UpdateStatus(ctx echo.Context, id int64, status int) (err error) {
// 	//TODO implement me
// 	panic("implement me")
// }

func New() domain.UserUsecase {
	return &userUsecase{}
}
