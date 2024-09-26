package entities

import "time"

type Post struct {
	Id        int64     `gorm:"column:id;primary_key'"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	Category  string    `gorm:"column:category"`
	Tags      string    `gorm:"column:tags"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (p *Post) TableName() string {
	return "posts"
}
