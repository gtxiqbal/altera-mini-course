package test

import (
	"fmt"
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/controller"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/repository"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/gtxiqbal/altera-mini-course/test/mock"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"strings"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)
	db = config.NewDBMySQL()
	userRepo = repository.NewUserRepositoryImpl(db)
	userSvc = service.NewUserServiceImpl(userRepo)
	userCtrl = controller.NewUserController(userSvc)
}

var (
	db           *gorm.DB                       = nil
	userRepo     *repository.UserRepositoryImpl = nil
	userSvc      *service.UserServiceImpl       = nil
	userCtrl     *controller.UserController     = nil
	echoMockUser                                = mock.EchoMock{E: echo.New()}
	userJson                                    = `{"name":"user_test","email":"user@gmail.com","password":"997891"}`
	tokenJwt                                    = `eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwOi8vbG9jYWxob3N0OjgwODAiLCJleHAiOjE2NjMzNzA5NDYsImlhdCI6MTY2MzM2NzM0NiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdDo4MDgwIiwianRpIjoiOGExNzc5ZTQtOWE1Ni00MzczLWI1MTItOGI4NzNmZDU4N2YzIiwibmJmIjoxNjYzMzY3MzQ2LCJzdWIiOiJ1c2VyX3Rlc3QiLCJ0eXAiOiJCZWFyZXIiLCJ1c2VyX2lkIjoiMTAifQ.rwLmvcOD_qAizbLTCWilm-prrGfTppfmAICQGcXxZNwDan877kpJrcCYTr8QPHhrbInz_0m6Z_Ez3bqc7WB_MA`
)

func TestUserGetAllSuccess(t *testing.T) {
	c, rec := echoMockUser.RequestMock(echo.GET, "/", nil)
	c.SetPath("/users")

	if assert.NoError(t, userCtrl.GetAll(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserGetByIDSuccess(t *testing.T) {
	c, rec := echoMockUser.RequestMock(echo.GET, "/", nil)
	c.SetPath("/users")
	c.SetParamNames("id")
	c.SetParamValues("9")

	//testing
	if assert.NotPanics(t, func() {
		_ = userCtrl.GetAll(c)
	}) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserGetByIDInvalid(t *testing.T) {
	c, _ := echoMockUser.RequestMock(echo.GET, "/", nil)
	c.SetPath("/users")
	c.SetParamNames("id")
	c.SetParamValues("6")

	//testing
	assert.Panics(t, func() {
		_ = userCtrl.GetByID(c)
	})
}

func TestUserCreateSuccess(t *testing.T) {
	c, rec := echoMockUser.RequestMock(echo.POST, "/", strings.NewReader(userJson))
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.SetPath("/users")

	//testing
	if assert.NotPanics(t, func() {
		_ = userCtrl.Create(c)
	}) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserCreateInvalid(t *testing.T) {
	c, _ := echoMockUser.RequestMock(echo.POST, "/", nil)
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.SetPath("/users")

	//testing
	assert.Panics(t, func() {
		_ = userCtrl.Create(c)
	})
}

func TestUserUpdateSuccess(t *testing.T) {
	c, rec := echoMockUser.RequestMock(echo.PUT, "/", strings.NewReader(userJson))
	c.Request().Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", tokenJwt))
	c.SetPath("/users")
	c.SetParamNames("id")
	c.SetParamValues("10")

	if assert.NotPanics(t, func() {
		_ = userCtrl.Update(c)
	}) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserUpdateInvalid(t *testing.T) {
	c, _ := echoMockUser.RequestMock(echo.PUT, "/", strings.NewReader(userJson))
	c.Request().Header.Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", tokenJwt))
	c.SetPath("/users")

	assert.Panics(t, func() {
		_ = userCtrl.Update(c)
	})
}

func TestUserDeleteSuccess(t *testing.T) {
	c, rec := echoMockUser.RequestMock(echo.DELETE, "/", nil)
	c.Request().Header.Add(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", tokenJwt))
	c.SetPath("/users")
	c.SetParamNames("id")
	c.SetParamValues("10")

	//testing
	if assert.NotPanics(t, func() {
		_ = userCtrl.Delete(c)
	}) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUserDeleteInvalid(t *testing.T) {
	c, _ := echoMockUser.RequestMock(echo.DELETE, "/", nil)
	c.SetPath("/users")
	c.SetParamNames("id")
	c.SetParamValues("9")

	//testing
	assert.Panics(t, func() {
		_ = userCtrl.Delete(c)
	})
}
