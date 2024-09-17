package entities

import "time"

type Article struct {
	Id          int64
	Title       string
	Content     string
	PublishDate time.Time
	Category    string
	Tags        []string
	UpdatedAt   time.Time
}
