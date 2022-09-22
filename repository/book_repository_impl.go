package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/gtxiqbal/altera-mini-course/helper"
	"github.com/gtxiqbal/altera-mini-course/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BookRepositoryImpl struct {
	collection *mongo.Collection
}

func NewBookRepositoryImpl(db *mongo.Database) *BookRepositoryImpl {
	collection := db.Collection("book")
	return &BookRepositoryImpl{collection: collection}
}

func (bookRepo *BookRepositoryImpl) FindAll(ctx context.Context) ([]model.Book, error) {
	csr, err := bookRepo.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var books []model.Book
	for csr.Next(ctx) {
		var book model.Book
		if err = csr.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (bookRepo *BookRepositoryImpl) FindByID(ctx context.Context, id string) (model.Book, error) {
	csr := bookRepo.collection.FindOne(ctx, bson.M{"id": id})
	var book model.Book
	helper.PanicIfError(csr.Decode(&book))
	return book, nil
}

func (bookRepo *BookRepositoryImpl) Insert(ctx context.Context, book *model.Book) error {
	book.ID = uuid.Must(uuid.NewRandom()).String()
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	_, err := bookRepo.collection.InsertOne(ctx, book)
	return err
}

func (bookRepo *BookRepositoryImpl) Update(ctx context.Context, book *model.Book) error {
	book.UpdatedAt = time.Now()
	_, err := bookRepo.collection.UpdateOne(ctx, bson.M{"id": book.ID}, bson.M{"$set": book})
	return err
}

func (bookRepo *BookRepositoryImpl) Delete(ctx context.Context, book *model.Book) error {
	_, err := bookRepo.collection.DeleteOne(ctx, bson.M{"id": book.ID})
	return err
}
