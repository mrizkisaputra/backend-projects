package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"personal-blog/internal/exceptions"
	"strconv"
)
import . "personal-blog/internal/models/entities"

type articleRepository struct {
	Log *logrus.Logger
}

func NewArticleRepository(log *logrus.Logger) ArticleRepositoryInterface {
	return &articleRepository{
		Log: log,
	}
}

func (self *articleRepository) isExist(id int64, articles []Article) (exist bool, index int) {
	for i, item := range articles {
		if id == item.Id {
			exist = true
			index = i
			break
		}
	}
	return exist, index
}

func (self *articleRepository) readDirectoryArticles(dir string) ([]os.DirEntry, error) {
	/* cari semua file .json didalam directory */
	files, errReadDir := os.ReadDir(dir)
	if errReadDir != nil {
		self.Log.Errorf("error reading directory %s", dir)
		return nil, errReadDir
	}
	return files, nil
}

func (self *articleRepository) WriteToFilesystem(article *Article, filename string, dir string) (bool, error) {
	if filepath.Ext(filename) != ".json" {
		self.Log.Error("filename must be JSON extension")
		return false, errors.New("file must have JSON extension")
	}

	fileLocation := filepath.Join(dir, filename)
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, errCreate := os.Create(fileLocation)
		if errCreate != nil {
			self.Log.Errorf("error creating file to: %s", fileLocation)
			return false, err
		}
		errEncoded := json.NewEncoder(file).Encode(article)
		if errEncoded != nil {
			self.Log.Error("error encoding article to json")
			return false, err
		}
		return true, nil
	} else {
		return false, errors.New(fmt.Sprintf("filename %s already exists", filename))
	}
}

func (self *articleRepository) ReadToFilesystem(dir string) ([]Article, error) {
	/* cari semua file .json didalam directory */
	files, err := self.readDirectoryArticles(dir)
	if err != nil {
		return nil, err
	}

	var articles []Article
	for _, file := range files {
		/* memastikan hanya file .json yang diambil */
		if filepath.Ext(file.Name()) == ".json" {
			/* baca isi file */
			data, errReadFile := os.ReadFile(filepath.Join(dir, file.Name()))
			if errReadFile != nil {
				self.Log.Errorf("error reading file %s", filepath.Join(dir, file.Name()))
				return nil, errReadFile
			}

			/* decode menjadi object struct article */
			var article Article
			errDecoded := json.Unmarshal(data, &article)
			if errDecoded != nil {
				self.Log.Errorf("error decoding file %s", filepath.Join(dir, file.Name()))
				return nil, errDecoded
			}
			articles = append(articles, article)
		}
	}
	return articles, nil
}

func (self *articleRepository) Create(article *Article, dir string) error {
	filename := strconv.FormatInt(article.Id, 10) + ".json"
	_, errWriteFile := self.WriteToFilesystem(article, filename, dir)
	if errWriteFile != nil {
		return errWriteFile
	}
	return nil
}

func (self *articleRepository) FindById(id int64, dir string) (*Article, error) {
	articles, errReadFilesystem := self.ReadToFilesystem(dir)
	if errReadFilesystem != nil {
		return new(Article), errReadFilesystem
	}

	if len(articles) < 1 {
		return nil, exceptions.NewErrNotFound("article is empty!")
	}

	exist, index := self.isExist(id, articles)
	if !exist {
		return new(Article), exceptions.NewErrNotFound(fmt.Sprintf("article with %d notfound!", id))
	}
	article := new(Article)
	article = &articles[index]
	return article, nil
}

func (self *articleRepository) FindAll(dir string) ([]Article, error) {
	articles, err := self.ReadToFilesystem(dir)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (self *articleRepository) Delete(id int64, dir string) (bool, error) {
	_, err := self.FindById(id, dir)
	if err != nil {
		return false, err
	}
	filename := filepath.Join(dir, fmt.Sprintf("%d.json", id))
	errRemove := os.Remove(filename)
	if errRemove != nil {
		return false, errRemove
	}
	return true, nil
}

func (self *articleRepository) Update(id int64, newArticle *Article, dir string) (bool, error) {
	//_, err := self.FindById(id, dir)
	//if err != nil {
	//	return false, err
	//}

	filename := filepath.Join(dir, fmt.Sprintf("%d.json", id))
	openFile, errOpenFile := os.OpenFile(filename, os.O_RDWR|os.O_TRUNC, 0644)
	if errOpenFile != nil {
		self.Log.Errorf("error open file %s", filename)
		return false, errOpenFile
	}

	errEncoded := json.NewEncoder(openFile).Encode(newArticle)
	if errEncoded != nil {
		return false, errEncoded
	}
	return true, nil
}
