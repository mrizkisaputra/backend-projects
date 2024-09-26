package dto

import "time"

type PostResponse struct {
	Id        int64
	Title     string
	Content   string
	Category  string
	Tags      []string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreatePostRequestBody struct {
	Title    string   `json:"title" validate:"required,min=3,max=150"`
	Content  string   `json:"content" validate:"required"`
	Category string   `json:"category" validate:"required,min=3,max=100,alpha"`
	Tags     []string `json:"tags"`
}

type UpdatePostRequestBody struct {
	Id       int64    `json:"-" validate:"required"`
	Title    string   `json:"title" validate:"required,min=1,max=150"`
	Content  string   `json:"content" validate:"required"`
	Category string   `json:"category" validate:"required,min=1,max=100"`
	Tags     []string `json:"tags"`
}

type GetPostRequestParam struct {
	Id int64 `json:"-" validate:"required"`
}

type DeletePostRequestParam struct {
	Id int64 `json:"-" validate:"required"`
}

type SearchPostRequestQueryParam struct {
	Title    string `json:"-" validate:"max=150"`
	Category string `json:"-" validate:"max=100"`
	Page     int    `json:"-" validate:"min=0"`
	Limit    int    `json:"-" validate:"max=100"`
}

type PostValidationErrResponse struct {
	Title    []string `json:"title,omitempty"`
	Content  []string `json:"content,omitempty"`
	Category []string `json:"category,omitempty"`
}
