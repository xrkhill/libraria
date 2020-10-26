package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/data"
	"github.com/xrkhill/libraria/internal/repository"
)

// BookService is an HTTP service
type BookService struct {
	repo *repository.BookRepository
}

// NewBookService returns a BookService struct
func NewBookService(r *repository.BookRepository) *BookService {
	return &BookService{
		repo: r,
	}
}

// AllBooks fetches all books from the collection
func (s *BookService) AllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, s.repo.ReadAll())
}

// GetBook fetches one book by ISBN
func (s *BookService) GetBook(c *gin.Context) {
	isbn := c.Params.ByName("isbn")

	book, err := s.repo.Read(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook creates one book
func (s *BookService) CreateBook(c *gin.Context) {
	var newBook data.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := s.repo.Create(newBook)
	if err != nil {
		c.JSON(http.StatusNotModified, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook creates one book
func (s *BookService) UpdateBook(c *gin.Context) {
	var existingBook data.Book
	if err := c.ShouldBindJSON(&existingBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := s.repo.Update(existingBook)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook creates one book
func (s *BookService) DeleteBook(c *gin.Context) {
	isbn := c.Params.ByName("isbn")

	if err := s.repo.Delete(isbn); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
