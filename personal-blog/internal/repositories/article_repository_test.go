package repositories

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"log"
	"personal-blog/internal/exceptions"
	"personal-blog/internal/models/entities"
	"testing"
	"time"
)

var articleRepo = NewArticleRepository(logrus.New())

func TestArticleRepository_WriteToFilesystem(t *testing.T) {
	tests := []struct {
		TestName string
		Article  entities.Article
		Filename string
		Dir      string
	}{
		{
			TestName: "[Test 1]: Should be return true and create file 1.json",
			Article:  entities.Article{Id: 1, Title: "Tutorial Java", Content: "java is programming language"},
			Filename: "1.json",
			Dir:      "../../articles/test",
		},
		{
			TestName: "[Test 2]: Should be return true and create file 2.json",
			Article:  entities.Article{Id: 2, Title: "Tutorial Golang", Content: "golang is programming language"},
			Filename: "2.json",
			Dir:      "../../articles/test",
		},
		{
			TestName: "[Test 3]: Should be return true and create file 3.json",
			Article:  entities.Article{Id: 3, Title: "Tutorial Kotlin", Content: "kotlin is programming language"},
			Filename: "3.json",
			Dir:      "../../articles/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			ok, err := articleRepo.WriteToFilesystem(&tt.Article, tt.Filename, tt.Dir)
			assert.Nil(t, err)
			assert.True(t, ok)
		})
	}
}

func TestArticleRepository_ReadToFilesystem(t *testing.T) {
	articles, err := articleRepo.ReadToFilesystem("../../articles/test")
	assert.Nil(t, err)
	assert.Equal(t, 5, len(articles))
	log.Print(articles)
}

func TestArticleRepository_Create(t *testing.T) {
	tests := []struct {
		TestName string
		Article  entities.Article
		Dir      string
	}{
		{
			TestName: "[Test 1]: Should be return not error",
			Article: entities.Article{
				Id:          time.Now().UnixNano(),
				Title:       "Tutorial Java",
				Content:     "java is programming language",
				PublishDate: time.Now().Local(),
			},
			Dir: "../../articles/test",
		},
		{
			TestName: "[Test 2]: Should be return not error",
			Article: entities.Article{
				Id:          time.Now().UnixNano() + 2,
				Title:       "Tutorial Golang",
				Content:     "golang is programming language",
				PublishDate: time.Now().Local(),
			},
			Dir: "../../articles/test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			err := articleRepo.Create(&tt.Article, tt.Dir)
			assert.Nil(t, err)
		})
	}
}

func TestArticleRepository_FindById(t *testing.T) {
	publishDate, _ := time.Parse(time.RFC3339Nano, "2024-09-14T11:27:24.8457157+07:00")
	updatedAt, _ := time.Parse(time.RFC3339, "0001-01-01T00:00:00Z")
	tests := []struct {
		TestName string
		Id       int64
		Dir      string
		Expected *entities.Article
	}{
		{
			TestName: "[Test 1]: Should be return article",
			Id:       1726288044845715700,
			Dir:      "../../articles/test",
			Expected: &entities.Article{
				Id:          1726288044845715700,
				Title:       "Tutorial Java",
				Content:     "java is programming language",
				PublishDate: publishDate,
				Tags:        nil,
				Category:    "",
				UpdatedAt:   updatedAt,
			},
		},
		{
			TestName: "[Test 2]: Should be return empty article",
			Id:       1726285956601761792,
			Dir:      "../../articles/test",
			Expected: new(entities.Article),
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			article, err := articleRepo.FindById(tt.Id, tt.Dir)
			assert.Nil(t, err)
			assert.Equal(t, tt.Expected, article)
			fmt.Println(article)
		})
	}

}

func TestArticleRepository_FindAll(t *testing.T) {
	articles, err := articleRepo.FindAll("../../articles/test")
	assert.Nil(t, err)
	assert.NotEmpty(t, articles)
	fmt.Println(articles)
}

func TestArticleRepository_Delete(t *testing.T) {
	tests := []struct {
		TestName string
		Id       int64
		Dir      string
		Expected bool
		Error    error
	}{
		{
			TestName: "[Test 1]: Should be return false, because ID notfound",
			Id:       1000000000000,
			Dir:      "../../articles/test",
			Expected: false,
			Error:    exceptions.NewErrNotFound("article with 1000000000000 notfound!"),
		},
		{
			TestName: "[Test 2]: Should be return true, delete successfully",
			Id:       1726307569307433502,
			Dir:      "../../articles/test",
			Expected: true,
			Error:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.TestName, func(t *testing.T) {
			ok, err := articleRepo.Delete(tt.Id, tt.Dir)
			assert.Equal(t, tt.Error, err)
			assert.Equal(t, tt.Expected, ok)
		})
	}

}

func TestArticleRepository_Update_Success(t *testing.T) {
	oldArticle, err := articleRepo.FindById(1, "../../articles/test")
	assert.Nil(t, err)

	newArticle := entities.Article{
		Id:          oldArticle.Id,
		Title:       oldArticle.Title,
		Content:     oldArticle.Content,
		PublishDate: oldArticle.PublishDate,
		Category:    oldArticle.Category,
		Tags:        []string{"programming", "coding", "programming language"},
		UpdatedAt:   time.Now().Local(),
	}

	ok, errUpdate := articleRepo.Update(1, &newArticle, "../../articles/test")
	assert.Nil(t, errUpdate)
	assert.True(t, ok)
}

func TestArticleRepository_Update_Fail(t *testing.T) {
	oldArticle, err := articleRepo.FindById(12345, "../../articles/test")
	assert.Error(t, err)

	newArticle := entities.Article{
		Id:          oldArticle.Id,
		Title:       oldArticle.Title,
		Content:     oldArticle.Content,
		PublishDate: oldArticle.PublishDate,
		Category:    oldArticle.Category,
		Tags:        []string{"programming", "coding", "programming language"},
		UpdatedAt:   time.Now().Local(),
	}

	ok, errUpdate := articleRepo.Update(12345, &newArticle, "../../articles/test")
	assert.Error(t, errUpdate)
	assert.False(t, ok)
}
