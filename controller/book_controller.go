package controller

import (
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/labstack/echo/v4"
)

type BookController struct {
	bookSvc service.BookService
}

func NewBookController(bookSvc service.BookService) *BookController {
	return &BookController{bookSvc: bookSvc}
}

func (bookCtrl *BookController) GetAll(c echo.Context) error {
	return c.JSON(200, bookCtrl.bookSvc.GetAll(c.Request().Context()))
}

func (bookCtrl *BookController) GetByID(c echo.Context) error {
	return c.JSON(200, bookCtrl.bookSvc.GetByID(c.Request().Context(), c.Param("id")))
}

func (bookCtrl *BookController) Create(c echo.Context) error {
	var bookDtoReq dto.BookDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&bookDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(bookDtoReq))
	return c.JSON(200, bookCtrl.bookSvc.Create(c.Request().Context(), bookDtoReq))
}

func (bookCtrl *BookController) Update(c echo.Context) error {
	var bookDtoReq dto.BookDtoReq
	helper.PanicIfErrorCode(400, c.Bind(&bookDtoReq))
	helper.PanicIfErrorCode(400, c.Validate(bookDtoReq))

	bookDtoReq.ID = c.Param("id")
	return c.JSON(200, bookCtrl.bookSvc.Update(c.Request().Context(), bookDtoReq))
}

func (bookCtrl *BookController) Delete(c echo.Context) error {
	return c.JSON(200, bookCtrl.bookSvc.Delete(c.Request().Context(), c.Param("id")))
}
