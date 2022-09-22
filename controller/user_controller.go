package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/labstack/echo/v4"
	"strconv"
	"strings"
)

type UserController struct {
	userSvc service.UserService
}

func NewUserController(userSvc service.UserService) *UserController {
	return &UserController{userSvc: userSvc}
}

func (userCtrl *UserController) GetAll(c echo.Context) error {
	return c.JSON(200, userCtrl.userSvc.GetAll(c.Request().Context()))
}

func (userCtrl *UserController) GetByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)
	return c.JSON(200, userCtrl.userSvc.GetByID(c.Request().Context(), uint(id)))
}

func (userCtrl *UserController) Create(c echo.Context) error {
	var userDtoReq dto.UserDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&userDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(userDtoReq))
	return c.JSON(200, userCtrl.userSvc.Create(c.Request().Context(), userDtoReq))
}

func (userCtrl *UserController) Update(c echo.Context) error {
	checkUser(c, "cannot delete another user")
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)

	var userDtoReq dto.UserDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&userDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(userDtoReq))

	userDtoReq.ID = uint(id)
	return c.JSON(200, userCtrl.userSvc.Update(c.Request().Context(), userDtoReq))
}

func (userCtrl *UserController) Delete(c echo.Context) error {
	checkUser(c, "cannot delete another user")
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfErrorCode(400, err)

	return c.JSON(200, userCtrl.userSvc.Delete(c.Request().Context(), uint(id)))
}

func checkUser(c echo.Context, message string) {
	var claims jwt.MapClaims
	token, ok := c.Get("user").(*jwt.Token)
	if ok {
		claims = token.Claims.(jwt.MapClaims)
	} else {
		_, claims, _ = config.JwtParse(strings.Replace(c.Request().Header.Get(echo.HeaderAuthorization), "Bearer ", "", -1))
	}
	if c.Param("id") != fmt.Sprint(claims["user_id"]) {
		helper.PanicErrorCode(400, errors.New(message))
	}
}
