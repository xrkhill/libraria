package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/data"
	"github.com/xrkhill/libraria/internal/repository"
	"github.com/xrkhill/libraria/internal/service"
)

func setupRouter(defaultBooks data.Books) *gin.Engine {
	repo := repository.NewBookRepository(defaultBooks)
	svc := service.NewBookService(repo)
	router := gin.Default()

	router.POST("/books", svc.CreateBook)
	router.GET("/books", svc.AllBooks)
	router.GET("/books/:isbn", svc.GetBook)
	router.PUT("/books", svc.UpdateBook)
	router.DELETE("/books/:isbn", svc.DeleteBook)

	return router
}

func main() {
	router := setupRouter(data.Books{})

	router.Run(":8080")
}
