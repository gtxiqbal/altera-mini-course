package repository

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model"
)

type BookRepository interface {
	FindAll(ctx context.Context) ([]model.Book, error)
	FindByID(ctx context.Context, id string) (model.Book, error)
	Insert(ctx context.Context, book *model.Book) error
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, book *model.Book) error
}
