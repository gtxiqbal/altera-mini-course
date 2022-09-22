package model

import (
	"time"
)

type Book struct {
	ID        string     `json:"id" bson:"id"`
	Title     string     `json:"title" bson:"title"`
	Isbn      string     `json:"isbn" bson:"isbn"`
	Writer    string     `json:"writer" bson:"writer"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}
