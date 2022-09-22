package middleware

import (
	"github.com/google/uuid"
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strings"
)

func RecoverConfig() middleware.RecoverConfig {
	return middleware.RecoverConfig{
		StackSize:    1 << 10,
		LogLevel:     log.ERROR,
		LogErrorFunc: helper.CustomHttpErrorHandler,
	}
}

func RequestIDConfig() middleware.RequestIDConfig {
	return middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
	}
}

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${id} | ${remote_ip} | ${host} | ${method} | ${path}\n",
	}
}

func JWTConfig() middleware.JWTConfig {
	return middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasPrefix(c.Request().RequestURI, "/books") && c.Request().Method == "GET" {
				return true
			}
			if (strings.HasPrefix(c.Request().RequestURI, "/users") ||
				strings.HasPrefix(c.Request().RequestURI, "/login")) &&
				c.Request().Method == "POST" {
				return true
			}
			return false
		},
		ParseTokenFunc: func(token string, c echo.Context) (interface{}, error) {
			tokenParse, claims, err := config.JwtParse(token)
			if err != nil {
				return nil, echo.NewHTTPError(401, err)
			}
			if claims["typ"] != string(service.Bearer) {
				return nil, echo.NewHTTPError(401, "token is not bearer")
			}
			return tokenParse, nil
		},
		ErrorHandlerWithContext: func(err error, c echo.Context) error {
			if err != nil {
				code := 401
				errorDesc := err.Error()
				httpError, ok := err.(*echo.HTTPError)
				if ok {
					code = httpError.Code
					err2, ok2 := httpError.Message.(error)
					if ok2 {
						errorDesc = err2.Error()
					} else {
						errorDesc = httpError.Message.(string)
					}
				}
				return c.JSON(code, echo.Map{
					"error":             "invalid_token",
					"error_description": errorDesc,
				})
			}
			return nil
		},
	}
}
