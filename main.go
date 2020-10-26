package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/data"
	"github.com/xrkhill/libraria/internal/repository"
	"github.com/xrkhill/libraria/internal/service"
)

func setupRouter(defaultBooks data.Books) *gin.Engine {
	repo := repository.NewMemoryBookRepository(defaultBooks)
	svc := service.NewBookService(repo)
	router := gin.Default()

	router.POST("/books", svc.Create)
	router.GET("/books", svc.ReadAll)
	router.GET("/books/:isbn", svc.ReadOne)
	router.PUT("/books", svc.Update)
	router.DELETE("/books/:isbn", svc.Delete)

	return router
}

func main() {
	router := setupRouter(data.Books{})

	router.Run(":8080")
}
