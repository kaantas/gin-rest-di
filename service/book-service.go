package service

import (
	"gin-rest-di/model"
	"gin-rest-di/repository"
	"github.com/gin-gonic/gin"
)

type IBookService interface {
	GetBookByNameAndAuthorName(context *gin.Context)
	GetAllBooks(context *gin.Context)
}

type BookService struct {
	injectedField string
}

func NewBookService(injectedField string) IBookService {
	return &BookService{injectedField: injectedField}
}

func (service *BookService) GetBookByNameAndAuthorName(context *gin.Context) {
	title := context.Query("title")
	authorFirstName := context.Query("firstName")

	bookList := repository.GetBooksList()
	foundIndex, foundValue := bookList.Find(func(index int, value interface{}) bool {
		book := value.(model.Book)
		return book.Title == title && book.Author.FirstName == authorFirstName
	})
	if foundIndex != -1 {
		context.JSON(200, foundValue)
	} else {
		context.JSON(404, gin.H{
			"message": service.injectedField,
		})
	}
}

func (service *BookService) GetAllBooks(context *gin.Context) {
	context.JSON(200, repository.GetBooks())
}
