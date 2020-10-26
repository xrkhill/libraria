package repository

import (
	"github.com/xrkhill/libraria/internal/data"
)

// BookRepository is a data mapper interface for the persistence layer
type BookRepository interface {
	Create(data.Book) (data.Book, error)
	ReadAll() data.Books
	Read(string) (data.Book, error)
	Update(data.Book) (data.Book, error)
	Delete(string) error
}
