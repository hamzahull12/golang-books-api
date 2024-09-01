package models

import "time"

type Book struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Author string `json:"author"`
	Summary string `json:"summary"`
	Publisher string `json:"publisher"`
	PageCount int `json:"pageCount"`
	ReadPage int `json:"readPage"`
	Finished bool`json:"finished"`
	Reading bool `json:"reading"`
	InsertedAt time.Time`json:"insertedAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateBookResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
			BookID string `json:"bookId"`
	} `json:"data"`
}
