package main

import (
	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/repository"
	"github.com/xrkhill/libraria/internal/service"
)

func main() {
	/*
		b := books.Books{
			"9780451524935": books.Book{
				Author:    "George Orwell",
				Title:     "1984",
				ISBN:      "9780451524935",
				Language:  "English",
				Published: time.Date(1961, time.January, 1, 0, 0, 0, 0, time.UTC),
				ListPrice: 999,
			},
			"9780553380163": books.Book{
				Author:    "Stephen Hawking",
				Title:     "A Brief History of Time",
				ISBN:      "9780553380163",
				Language:  "English",
				Published: time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
				ListPrice: 1800,
			},
		}
	//*/
	//b := books.Books{}
	repo := repository.NewBookRepository()
	svc := service.NewService(repo)

	route := gin.Default()
	route.GET("/books", svc.AllBooks)
	route.GET("/books/:isbn", svc.GetBook)
	route.POST("/books", svc.CreateBook)
	route.PUT("/books/:isbn", svc.UpdateBook)
	route.DELETE("/books/:isbn", svc.DeleteBook)
	route.Run(":8080")
}
