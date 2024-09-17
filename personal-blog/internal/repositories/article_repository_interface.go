package repositories

import . "personal-blog/internal/models/entities"

type ArticleRepositoryInterface interface {
	WriteToFilesystem(article *Article, filename string, dir string) (bool, error)

	ReadToFilesystem(dir string) ([]Article, error)

	Create(article *Article, dir string) error

	FindById(id int64, dir string) (*Article, error)

	FindAll(dir string) ([]Article, error)

	Delete(id int64, dir string) (bool, error)

	Update(id int64, newArticle *Article, dir string) (bool, error)
}
