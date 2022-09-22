package service

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
)

type UserService interface {
	GetAll(ctx context.Context) dto.ResponseDTO[[]dto.UserDtoRes]
	GetByID(ctx context.Context, id uint) dto.ResponseDTO[dto.UserDtoRes]
	Create(ctx context.Context, userDto dto.UserDtoReq) dto.ResponseDTO[dto.UserDtoRes]
	Update(ctx context.Context, userDtoReq dto.UserDtoReq) dto.ResponseDTO[dto.UserDtoRes]
	Delete(ctx context.Context, id uint) dto.ResponseDTO[any]
}
