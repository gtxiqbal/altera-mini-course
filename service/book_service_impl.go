package service

import (
	"context"
	"errors"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model"
	"github.com/gtxiqbal/altera-mini-course/model/dto"
	"github.com/gtxiqbal/altera-mini-course/repository"
)

type BookServiceImpl struct {
	bookRepo repository.BookRepository
}

func NewBookServiceImpl(bookRepo repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{bookRepo: bookRepo}
}

func (bookSvc *BookServiceImpl) GetAll(ctx context.Context) dto.ResponseDTO[[]dto.BookDtoRes] {
	books, err := bookSvc.bookRepo.FindAll(ctx)
	helper.PanicIfErrorCode(400, err)
	var booksDtoRes []dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&booksDtoRes, books))
	return dto.ResponseDTO[[]dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get all book",
		Data:    booksDtoRes,
	}
}

func (bookSvc *BookServiceImpl) getByID(ctx context.Context, id string) model.Book {
	book, err := bookSvc.bookRepo.FindByID(ctx, id)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			helper.PanicErrorCode(404, errors.New("record not found"))
		}
		helper.PanicErrorCode(404, err)
	}
	return book
}

func (bookSvc *BookServiceImpl) GetByID(ctx context.Context, id string) dto.ResponseDTO[dto.BookDtoRes] {
	book := bookSvc.getByID(ctx, id)

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success get book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Create(ctx context.Context, bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes] {
	var book model.Book
	helper.PanicIfError(helper.MappingStruct(&book, bookDtoReq))
	helper.PanicIfErrorCode(404, bookSvc.bookRepo.Insert(ctx, &book))

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success create book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Update(ctx context.Context, bookDtoReq dto.BookDtoReq) dto.ResponseDTO[dto.BookDtoRes] {
	b := bookSvc.getByID(ctx, bookDtoReq.ID)
	var book model.Book
	helper.PanicIfError(helper.MappingStruct(&book, bookDtoReq))
	book.CreatedAt = b.CreatedAt
	helper.PanicIfErrorCode(400, bookSvc.bookRepo.Update(ctx, &book))

	var bookDtoRes dto.BookDtoRes
	helper.PanicIfError(helper.MappingStruct(&bookDtoRes, book))

	return dto.ResponseDTO[dto.BookDtoRes]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success update book",
		Data:    bookDtoRes,
	}
}

func (bookSvc *BookServiceImpl) Delete(ctx context.Context, id string) dto.ResponseDTO[any] {
	book := bookSvc.getByID(ctx, id)
	helper.PanicIfErrorCode(400, bookSvc.bookRepo.Delete(ctx, &book))
	return dto.ResponseDTO[any]{
		Code:    200,
		Status:  dto.StatusSuccess,
		Message: "success delete book",
	}
}
