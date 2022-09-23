package router

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gtxiqbal/altera-mini-course/controller"
	"github.com/gtxiqbal/altera-mini-course/helper"
	middleware2 "github.com/gtxiqbal/altera-mini-course/middleware"
	"github.com/gtxiqbal/altera-mini-course/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

type Router struct {
	e *echo.Echo
}

func NewRouter() *Router {
	e := echo.New()

	e.IPExtractor = echo.ExtractIPFromXFFHeader()
	e.Validator = &helper.CustomValidator{Validator: validator.New()}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.RecoverWithConfig(middleware2.RecoverConfig()))
	e.Use(middleware.RequestIDWithConfig(middleware2.RequestIDConfig()))
	e.Use(middleware.LoggerWithConfig(middleware2.LoggerConfig()))
	e.Use(middleware.JWTWithConfig(middleware2.JWTConfig()))
	return &Router{e: e}
}

func (r *Router) StartServer() {
	r.e.Logger.Fatal(r.e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func (r *Router) SetRouteUser(userService service.UserService) {
	userController := controller.NewUserController(userService)
	groupUser := r.e.Group("/users")
	groupUser.GET("", userController.GetAll)
	groupUser.GET("/:id", userController.GetByID)
	groupUser.POST("", userController.Create)
	groupUser.PUT("/:id", userController.Update)
	groupUser.DELETE("/:id", userController.Delete)
}

func (r *Router) SetRouteBook(bookService service.BookService) {
	bookController := controller.NewBookController(bookService)
	groupBook := r.e.Group("/books")
	groupBook.GET("", bookController.GetAll)
	groupBook.GET("/:id", bookController.GetByID)
	groupBook.POST("", bookController.Create)
	groupBook.PUT("/:id", bookController.Update)
	groupBook.DELETE("/:id", bookController.Delete)
}

func (r *Router) SetRouteLogin(loginService service.LoginService) {
	loginController := controller.NewLoginController(loginService)
	groupLogin := r.e.Group("/login")
	groupLogin.POST("", loginController.DoLogin)
	groupLogin.POST("/refresh", loginController.DoRefreshToken)
}
