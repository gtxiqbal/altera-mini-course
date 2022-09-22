package service

import (
	"context"
	"errors"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserServiceImpl(userRepo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepo: userRepo}
}

func (userSvc *UserServiceImpl) GetAll(ctx context.Context) dto.ResponseDTO[[]dto.UserDtoRes] {
	users, err := userSvc.userRepo.FindAll(ctx)
	helper.PanicIfError(err)

	var usersDtoRes []dto.UserDtoRes
	helper.PanicIfError(helper.MappingStruct(&usersDtoRes, users))

	return dto.ResponseDTO[[]dto.UserDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get all users",
		Data:    usersDtoRes,
	}
}

func (userSvc *UserServiceImpl) getByID(ctx context.Context, id uint) model.User {
	user, err := userSvc.userRepo.FindByID(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			helper.PanicErrorCode(404, errors.New("user not found"))
		}
		helper.PanicErrorCode(404, err)
	}
	return user
}

func (userSvc *UserServiceImpl) GetByID(ctx context.Context, id uint) dto.ResponseDTO[dto.UserDtoRes] {
	user := userSvc.getByID(ctx, id)
	var userDtoRes dto.UserDtoRes
	helper.PanicIfError(helper.MappingStruct(&userDtoRes, user))

	return dto.ResponseDTO[dto.UserDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get user",
		Data:    userDtoRes,
	}
}

func (userSvc *UserServiceImpl) Create(ctx context.Context, userDtoReq dto.UserDtoReq) dto.ResponseDTO[dto.UserDtoRes] {
	password, err := bcrypt.GenerateFromPassword([]byte(userDtoReq.Password), bcrypt.DefaultCost)
	helper.PanicIfErrorCode(400, err)
	userDtoReq.Password = string(password)

	var user model.User
	helper.PanicIfError(helper.MappingStruct(&user, userDtoReq))
	helper.PanicIfError(userSvc.userRepo.Insert(ctx, &user))

	var userDtoRes dto.UserDtoRes
	helper.PanicIfError(helper.MappingStruct(&userDtoRes, user))

	return dto.ResponseDTO[dto.UserDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success create user",
		Data:    userDtoRes,
	}
}

func (userSvc *UserServiceImpl) Update(ctx context.Context, userDtoReq dto.UserDtoReq) dto.ResponseDTO[dto.UserDtoRes] {
	password, err := bcrypt.GenerateFromPassword([]byte(userDtoReq.Password), bcrypt.DefaultCost)
	helper.PanicIfErrorCode(400, err)
	userDtoReq.Password = string(password)

	u := userSvc.getByID(ctx, userDtoReq.ID)
	var user model.User
	helper.PanicIfError(helper.MappingStruct(&user, userDtoReq))

	user.CreatedAt = u.CreatedAt
	helper.PanicIfError(userSvc.userRepo.Update(ctx, &user))

	var userDtoRes dto.UserDtoRes
	helper.PanicIfError(helper.MappingStruct(&userDtoRes, user))

	return dto.ResponseDTO[dto.UserDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success update user",
		Data:    userDtoRes,
	}
}

func (userSvc *UserServiceImpl) Delete(ctx context.Context, id uint) dto.ResponseDTO[any] {
	user := userSvc.getByID(ctx, id)
	helper.PanicIfError(userSvc.userRepo.Delete(ctx, &user))
	return dto.ResponseDTO[any]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success delete user",
	}
}
