package dto

import "time"

type BookDtoReq struct {
	ID     string `json:"id"`
	Title  string `json:"title" validate:"required"`
	Isbn   string `json:"isbn" validate:"required"`
	Writer string `json:"writer" validate:"required"`
}

type BookDtoRes struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Isbn      string    `json:"isbn"`
	Writer    string    `json:"writer"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
