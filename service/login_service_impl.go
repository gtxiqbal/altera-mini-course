package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/repository"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
	"time"
)

type TokenType string

var (
	Bearer  TokenType = "Bearer"
	Refresh TokenType = "Refresh"
)

type LoginServiceImpl struct {
	userRepo repository.UserRepository
}

func NewLoginServiceImpl(userRepo repository.UserRepository) *LoginServiceImpl {
	return &LoginServiceImpl{userRepo: userRepo}
}

func (loginSvc *LoginServiceImpl) DoLogin(ctx context.Context, loginDtoReq dto.LoginDtoReq) map[string]any {
	user, err := loginSvc.userRepo.FindByEmail(ctx, loginDtoReq.Email)
	if err != nil {
		if err.Error() == "record not found" {
			helper.PanicErrorCode(404, errors.New("user not found"))
		}
		helper.PanicErrorCode(400, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDtoReq.Password))
	if err != nil {
		helper.PanicErrorCode(400, errors.New("password doesn't match"))
	}

	return generateToken(user)
}

func (loginSvc *LoginServiceImpl) DoRefreshToken(ctx context.Context, refreshDtoReq dto.RefreshDtoReq) map[string]any {
	if refreshDtoReq.GrantType != dto.RefreshType {
		helper.PanicErrorCode(400, errors.New("grant_type must refresh_token"))
	}

	_, claims, err := config.JwtParse(refreshDtoReq.RefreshToken)
	helper.PanicIfErrorCode(400, err)

	user, err := loginSvc.userRepo.FindByEmail(ctx, fmt.Sprint(claims["sub"]))
	if err != nil {
		if err.Error() == "record not found" {
			helper.PanicErrorCode(404, errors.New("user not found"))
		}
		helper.PanicErrorCode(400, err)
	}

	return generateToken(user)
}

func generateToken(user model.User) map[string]any {
	//accessClaims access_token
	accessClaims := jwt.MapClaims{}
	accessClaims["aud"] = fmt.Sprintf("%s:%s", os.Getenv("jwt_issuer"), os.Getenv("SERVER_PORT"))
	accessClaims["exp"] = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	accessClaims["jti"] = uuid.Must(uuid.NewRandom()).String()
	accessClaims["iat"] = jwt.NewNumericDate(time.Now())
	accessClaims["iss"] = accessClaims["aud"]
	accessClaims["nbf"] = accessClaims["iat"]
	accessClaims["sub"] = user.Name
	accessClaims["typ"] = Bearer
	accessClaims["user_id"] = strconv.Itoa(int(user.ID))

	//generate access_token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims)
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	helper.PanicIfErrorCode(400, err)

	//refreshClaims access_token
	refreshClaims := jwt.MapClaims{}
	refreshClaims["aud"] = accessClaims["aud"]
	refreshClaims["exp"] = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	refreshClaims["jti"] = uuid.Must(uuid.NewRandom()).String()
	refreshClaims["iat"] = jwt.NewNumericDate(time.Now())
	refreshClaims["iss"] = refreshClaims["aud"]
	refreshClaims["nbf"] = refreshClaims["iat"]
	refreshClaims["sub"] = user.Email
	refreshClaims["typ"] = Refresh

	//generate refresh_token
	token = jwt.NewWithClaims(jwt.SigningMethodHS512, refreshClaims)
	refreshToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	helper.PanicIfErrorCode(400, err)

	return map[string]any{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
}
