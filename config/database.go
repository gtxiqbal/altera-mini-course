package config

import (
	"context"
	"fmt"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strconv"
	"time"
)

func NewDBMySQL() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	dbIdleConn, _ := strconv.Atoi(os.Getenv("DB_IDLE_CONN"))
	dbOpenConn, _ := strconv.Atoi(os.Getenv("DB_OPEN_CONN"))
	dbMaxIdleTimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONN"))
	dbMaxLifeTimeConn, _ := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN"))

	sqlDB.SetMaxIdleConns(dbIdleConn)
	sqlDB.SetMaxOpenConns(dbOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbMaxIdleTimeConn) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxLifeTimeConn) * time.Minute)

	isMigrate, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	helper.PanicIfError(err)
	if isMigrate {
		helper.PanicIfError(autoMigrate(db))
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&model.User{})
}

func NewDBMongo() *mongo.Database {
	clientOptions := options.Client()
	clientOptions.ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority",
		os.Getenv("DB_USERNAME_MONGO"),
		os.Getenv("DB_PASSWORD_MONGO"),
		os.Getenv("DB_HOST_MONGO")))
	connect, err := mongo.Connect(context.Background(), clientOptions)
	helper.PanicIfError(err)
	db := connect.Database(os.Getenv("DB_NAME"))
	isMigrate, err := strconv.ParseBool(os.Getenv("AUTO_MIGRATE"))
	helper.PanicIfError(err)
	if isMigrate {
		helper.PanicIfError(db.Collection("book").Drop(context.Background()))
		helper.PanicIfError(db.CreateCollection(context.Background(), "book"))
	}
	return db
}
