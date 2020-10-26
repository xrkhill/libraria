package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/repository"
	"github.com/xrkhill/libraria/internal/service"
)

func main() {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)
	route := gin.Default()

	route.GET("/books", svc.AllBooks)
	route.GET("/books/:isbn", svc.GetBook)
	route.POST("/books", svc.CreateBook)
	route.PUT("/books", svc.UpdateBook)
	route.DELETE("/books/:isbn", svc.DeleteBook)

	route.Run(":8080")
}
