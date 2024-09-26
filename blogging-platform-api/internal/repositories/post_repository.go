package repositories

import (
	"blogging-platform-api/internal/models/dto"
	"blogging-platform-api/internal/models/entities"
	"gorm.io/gorm"
)

type postRepository struct {
}

func NewPosRepository() PostRepositoryInterface {
	repo := new(postRepository)
	return repo
}

func (p *postRepository) Create(db *gorm.DB, entity *entities.Post) (*entities.Post, error) {
	/**
	SQL: INSERT INTO posts (id,title,content,category,tags,created_at,updated_at
		 VALUES (... ... ... ... ... ... ...)
	*/
	if err := db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (p *postRepository) FindById(db *gorm.DB, id int64) (*entities.Post, error) {
	var post entities.Post
	/**
	SQL: SELECT * FROM posts WHERE id = ? LIMIT 1
	*/
	if err := db.Model(entities.Post{}).Where("id = ?", id).Take(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *postRepository) Update(db *gorm.DB, entity *entities.Post) (*entities.Post, error) {
	/**
	SQL: UPDATE posts SET title = ? content = ? category = ? tags = ?
	*/
	if err := db.Save(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (p *postRepository) Remove(db *gorm.DB, id int64) error {
	/**
	SQL: DELETE FROM posts WHERE id = ?
	*/
	return db.Where("id = ?", id).Delete(entities.Post{}).Error
}

func (p *postRepository) FindAll(db *gorm.DB, request *dto.SearchPostRequestQueryParam) (*[]entities.Post, int64, error) {
	var posts []entities.Post

	/**
	kondisi pertama
	SQL: SELECT * FROM posts LIMIT 0,10

	kondisi kedua
	SQL: SELECT * FROM posts WHERE title LIKE ? LIMIT 0,10

	kondisi ketiga
	SQL: SELECT * FROM posts WHERE category LIKE ? LIMIT 0,10
	*/
	err := db.Model(entities.Post{}).
		Scopes(p.filter(request)).
		Offset((request.Page - 1) * request.Limit).
		Limit(request.Limit).
		Find(&posts).Error

	if err != nil {
		return nil, 0, err
	}

	/**
	menghitung total query post
	SQL: SELECT count(*) FROM posts ?

	menghitung total query post dengan kondisi title
	SQL: SELECT count(*) FROM posts WHERE title LIKE ?

	menghitung total query post dengan kondisi category
	SQL: SELECT count(*) FROM posts WHERE category LIKE ?
	*/
	var total int64
	if err := db.Model(entities.Post{}).Scopes(p.filter(request)).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return &posts, total, nil
}

type filterHandle func(db *gorm.DB) *gorm.DB

func (p *postRepository) filter(request *dto.SearchPostRequestQueryParam) filterHandle {
	return func(db *gorm.DB) *gorm.DB {
		if request.Title != "" {
			t := "%" + request.Title + "%"
			db.Where("title LIKE ?", t)
		}

		if request.Category != "" {
			t := "%" + request.Category + "%"
			db.Where("category LIKE ?", t)
		}
		return db
	}
}
