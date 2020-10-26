package repository

import (
	"fmt"
	"sync"

	"github.com/xrkhill/libraria/internal/data"
)

// BookRepository maps book structs to storage
type BookRepository struct {
	books data.Books
	mu    sync.RWMutex
}

// NewBookRepository rerurns a reference to a BookRepository
func NewBookRepository(defaultBooks data.Books) *BookRepository {
	return &BookRepository{
		books: defaultBooks,
	}
}

// Create creates a new book
func (b *BookRepository) Create(book data.Book) (data.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.books[book.ISBN]; ok {
		return data.Book{}, fmt.Errorf("book %s already exists, unable to create", book.ISBN)
	}

	b.books[book.ISBN] = book

	return book, nil
}

// ReadAll returns a list of books
func (b *BookRepository) ReadAll() data.Books {
	b.mu.RLock()
	defer b.mu.RUnlock()

	return b.books
}

func (b *BookRepository) Read(isbn string) (data.Book, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if existingBook, ok := b.books[isbn]; ok {
		return existingBook, nil
	}

	return data.Book{}, fmt.Errorf("book %s does not exist, unable to read", isbn)
}

// Update updates an existing book
func (b *BookRepository) Update(book data.Book) (data.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.books[book.ISBN]; ok {
		b.books[book.ISBN] = book
		return book, nil
	}

	return data.Book{}, fmt.Errorf("book %s does not exist, unable to update", book.ISBN)
}

// Delete removes an existing book
func (b *BookRepository) Delete(isbn string) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.books[isbn]; ok {
		delete(b.books, isbn)
		return nil
	}

	return fmt.Errorf("book %s does not exist, unable to delete", isbn)
}
