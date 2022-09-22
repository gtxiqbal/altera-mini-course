package main

import (
	"github.com/gtxiqbal/altera-mini-course/config"
	"github.com/gtxiqbal/altera-mini-course/repository"
	"github.com/gtxiqbal/altera-mini-course/router"
	"github.com/gtxiqbal/altera-mini-course/service"
)

func main() {
	//err := godotenv.Load()
	//helper.PanicIfError(err)

	db := config.NewDBMySQL()
	mongoDb := config.NewDBMongo()
	userRepositoryImpl := repository.NewUserRepositoryImpl(db)
	userServiceImpl := service.NewUserServiceImpl(userRepositoryImpl)

	bookRepositoryImpl := repository.NewBookRepositoryImpl(mongoDb)
	bookServiceImpl := service.NewBookServiceImpl(bookRepositoryImpl)

	loginServiceImpl := service.NewLoginServiceImpl(userRepositoryImpl)

	r := router.NewRouter()
	r.SetRouteUser(userServiceImpl)
	r.SetRouteBook(bookServiceImpl)
	r.SetRouteLogin(loginServiceImpl)
	r.StartServer()
}
