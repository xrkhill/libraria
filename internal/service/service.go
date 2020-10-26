package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xrkhill/libraria/internal/data"
	"github.com/xrkhill/libraria/internal/repository"
)

// Service is an HTTP service
type Service struct {
	repo *repository.BookRepository
}

// NewService returns a Service struct
func NewService(r *repository.BookRepository) *Service {
	return &Service{
		repo: r,
	}
}

// AllBooks fetches all books from the collection
func (s *Service) AllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, s.repo.ReadAll())
}

// GetBook fetches one book by ISBN
func (s *Service) GetBook(c *gin.Context) {
	isbn := c.Params.ByName("isbn")

	book, err := s.repo.Read(isbn)
	if err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusOK, book)
}

// CreateBook creates one book
func (s *Service) CreateBook(c *gin.Context) {
	var newBook data.Book
	if c.Bind(&newBook) != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	book, err := s.repo.Create(newBook)
	if err != nil {
		c.JSON(http.StatusNotModified, nil)
		return
	}

	c.JSON(http.StatusCreated, book)
}

// UpdateBook creates one book
func (s *Service) UpdateBook(c *gin.Context) {
}

// DeleteBook creates one book
func (s *Service) DeleteBook(c *gin.Context) {
	isbn := c.Params.ByName("isbn")

	if err := s.repo.Delete(isbn); err != nil {
		c.JSON(http.StatusNotFound, nil)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
