package usecase

import (
	"github.com/stretchr/testify/mock"
	. "personal-blog/internal/models/entities"
)

type MockArticleRepository struct {
	Mock mock.Mock
}

func (self *MockArticleRepository) WriteToFilesystem(article *Article, filename string, dir string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (self *MockArticleRepository) ReadToFilesystem(dir string) ([]Article, error) {
	//TODO implement me
	panic("implement me")
}

func (self *MockArticleRepository) Create(article *Article, dir string) error {
	args := self.Mock.Called(article, dir)
	return args.Error(0) // mengambil hasil simulasi pemanggilan
}

func (self *MockArticleRepository) FindById(id int64, dir string) (*Article, error) {
	args := self.Mock.Called(id, dir)
	article := args.Get(0).(Article)
	return &article, args.Error(1)
}

func (self *MockArticleRepository) FindAll(dir string) ([]Article, error) {
	//TODO implement me
	panic("implement me")
}

func (self *MockArticleRepository) Delete(id int64, dir string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (self *MockArticleRepository) Update(id int64, newArticle *Article, dir string) (bool, error) {
	//TODO implement me
	panic("implement me")
}
