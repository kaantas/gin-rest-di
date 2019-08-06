package controller

import (
	"gin-rest-di/service"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	BookService service.IBookService
}

func NewBookController(bookService service.IBookService) Controller {
	return &BookController{BookService: bookService}
}

func (controller BookController) Inject(engine *gin.Engine) {
	engine.GET("/query", controller.BookService.GetBookByNameAndAuthorName)
	engine.GET("/", controller.BookService.GetAllBooks)
}
