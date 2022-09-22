package service

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
)

type BookService interface {
	GetAll(ctx context.Context) dto.ResponseDTO[[]dto.BookDtoRes]
	GetByID(ctx context.Context, id string) dto.ResponseDTO[dto.BookDtoRes]
	Create(ctx context.Context, bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes]
	Update(ctx context.Context, bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes]
	Delete(ctx context.Context, id string) dto.ResponseDTO[any]
}
