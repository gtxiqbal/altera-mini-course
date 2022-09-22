package helper

import (
	"fmt"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/labstack/echo/v4"
)

func PanicIfError(err error) {
	PanicIfErrorCode(500, err)
}

func PanicIfErrorCode(code int, err error) {
	if err != nil {
		PanicErrorCode(code, err)
	}
}

func PanicErrorCode(code int, err error) {
	err = echo.NewHTTPError(code, err)
	panic(err)
}

func CustomHttpErrorHandler(c echo.Context, err error, stack []byte) error {
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		httpError = &echo.HTTPError{
			Code:     500,
			Message:  err.Error(),
			Internal: err,
		}
	}

	c.Logger().Error(httpError.Message)
	fmt.Println(string(stack))

	fieldErrorResponses, ok := httpError.Message.(*FieldErrorResponses)
	if ok {
		return c.JSON(httpError.Code, dto.ResponseDTO[any]{
			Code:    httpError.Code,
			Status:  dto.StatusFailed,
			Message: fieldErrorResponses.Error(),
			Error:   fieldErrorResponses,
		})
	}

	message := ""
	if err, ok = httpError.Message.(error); ok {
		message = err.Error()
	} else {
		message = httpError.Message.(string)
	}

	return c.JSON(httpError.Code, dto.ResponseDTO[any]{
		Code:    httpError.Code,
		Status:  dto.StatusFailed,
		Message: message,
	})
}
