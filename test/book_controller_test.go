package test

import (
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/controller"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/repository"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/gtxiqbal/altera-mini-course/test/mock"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	helper.PanicIfError(err)
	mongoDB := config.NewDBMongo()
	bookRepo = repository.NewBookRepositoryImpl(mongoDB)
	bookSvc = service.NewBookServiceImpl(bookRepo)
	bookCtrl = controller.NewBookController(bookSvc)
}

var (
	bookCtrl     *controller.BookController     = nil
	bookSvc      *service.BookServiceImpl       = nil
	bookRepo     *repository.BookRepositoryImpl = nil
	echoMockBook                                = mock.EchoMock{E: echo.New()}
	bookJson                                    = `{"title":"contoh buku 2", "isbn":"887697-271731-123","writer":"iqbal"}`
)

func TestGetAllSuccess(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.GET, "/", nil)
	c.SetPath("/books")

	//testing
	if assert.NoError(t, bookCtrl.GetAll(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestGetAllInvalid(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.POST, "/books", nil)

	//testing
	echoMockBook.E.ServeHTTP(rec, c.Request())
	assert.Equal(t, 400, rec.Code)
}

func TestGetByIDSuccess(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.GET, "/books", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	//testing
	if assert.NoError(t, bookCtrl.GetByID(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestGetByIDInvalid(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.GET, "/books", nil)
	c.SetParamNames("id")
	c.SetParamValues("2")

	//testing
	echoMockBook.E.ServeHTTP(rec, c.Request())
	assert.Equal(t, 404, rec.Code)
}

func TestCreateBookSuccess(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.POST, "/books", strings.NewReader(bookJson))
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//testing
	if assert.NoError(t, bookCtrl.Create(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestCreateBookInvalid(t *testing.T) {
	c, _ := echoMockBook.RequestMock(echo.POST, "/books", nil)
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//testing
	assert.Panics(t, func() {
		_ = bookCtrl.Create(c)
	})
}

func TestUpdateBookSuccess(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.PUT, "/books", strings.NewReader(bookJson))
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.SetParamNames("id")
	c.SetParamValues("1")

	//testing
	if assert.NoError(t, bookCtrl.Update(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestUpdateBookInvalid(t *testing.T) {
	c, _ := echoMockBook.RequestMock(echo.PUT, "/books", strings.NewReader(bookJson))
	c.Request().Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	//testing
	assert.Panics(t, func() {
		_ = bookCtrl.Update(c)
	})
}

func TestDeleteBookSuccess(t *testing.T) {
	c, rec := echoMockBook.RequestMock(echo.DELETE, "/books", nil)
	c.SetParamNames("id")
	c.SetParamValues("1")

	//testing
	if assert.NoError(t, bookCtrl.Delete(c)) {
		assert.Equal(t, 200, rec.Code)
	}
}

func TestDeleteBookInvalid(t *testing.T) {
	c, _ := echoMockBook.RequestMock(echo.DELETE, "/books", nil)

	//testing
	assert.Panics(t, func() {
		_ = bookCtrl.Delete(c)
	})
}
