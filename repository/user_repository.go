package repository

import (
	"context"
	"github.com/gtxiqbal/altera-mini-course/model"
)

type UserRepository interface {
	FindAll(ctx context.Context) ([]model.User, error)
	FindByEmail(ctx context.Context, email string) (model.User, error)
	FindByID(ctx context.Context, id uint) (model.User, error)
	Insert(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
