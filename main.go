package main

import (
	"gin-rest-di/controller"
	"gin-rest-di/service"
	"github.com/gin-gonic/gin"
)

func main()  {
	//routing
	engine := gin.Default()

	//initialize controllers & services
	bookService := service.NewBookService("Book Not Found !")
	bookController := controller.NewBookController(bookService)

	bookController.Inject(engine)

	//start
	engine.Run()
}
