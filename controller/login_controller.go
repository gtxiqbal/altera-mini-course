package controller

import (
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/labstack/echo/v4"
)

type LoginController struct {
	loginSvc service.LoginService
}

func NewLoginController(loginSvc service.LoginService) *LoginController {
	return &LoginController{loginSvc: loginSvc}
}

func (loginCtrl *LoginController) DoLogin(c echo.Context) error {
	var loginDtoReq dto.LoginDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&loginDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(loginDtoReq))
	return c.JSON(200, loginCtrl.loginSvc.DoLogin(c.Request().Context(), loginDtoReq))
}

func (loginCtrl *LoginController) DoRefreshToken(c echo.Context) error {
	var refreshDtoReq dto.RefreshDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&refreshDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(refreshDtoReq))
	return c.JSON(200, loginCtrl.loginSvc.DoRefreshToken(c.Request().Context(), refreshDtoReq))
}
