package mock

import (
	"github.com/go-playground/validator/v10"
	"github.com/gtxiqbal/altera-mini-course/helper"
	middleware2 "github.com/gtxiqbal/altera-mini-course/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io"
	"net/http/httptest"
)

type EchoMock struct {
	E *echo.Echo
}

func (em *EchoMock) RequestMock(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	em.E.Validator = &helper.CustomValidator{Validator: validator.New()}
	em.E.Use(middleware.RecoverWithConfig(middleware2.RecoverConfig()))
	em.E.Use(middleware.JWTWithConfig(middleware2.JWTConfig()))
	req := httptest.NewRequest(method, path, body)
	rec := httptest.NewRecorder()
	c := em.E.NewContext(req, rec)
	return c, rec
}
