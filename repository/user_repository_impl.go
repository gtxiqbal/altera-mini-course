package repository

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	*gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (userRepo *UserRepositoryImpl) FindAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	err := userRepo.DB.WithContext(ctx).Find(&users).Error
	return users, err
}

func (userRepo *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := userRepo.DB.WithContext(ctx).Where(&model.User{Email: email}).First(&user).Error
	return user, err
}

func (userRepo *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (model.User, error) {
	var user model.User
	err := userRepo.DB.WithContext(ctx).First(&user, id).Error
	return user, err
}

func (userRepo *UserRepositoryImpl) Insert(ctx context.Context, user *model.User) error {
	return userRepo.DB.WithContext(ctx).Create(user).Error
}

func (userRepo *UserRepositoryImpl) Update(ctx context.Context, user *model.User) error {
	return userRepo.DB.WithContext(ctx).Save(user).Error
}

func (userRepo *UserRepositoryImpl) Delete(ctx context.Context, user *model.User) error {
	return userRepo.DB.WithContext(ctx).Delete(user, user.ID).Error
}
