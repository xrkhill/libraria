package data

import (
	"time"
)

// Book is a struct for library book data
type Book struct {
	Author    string    `json:"author" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	ISBN      string    `json:"ISBN" binding:"required"`
	Language  string    `json:"language"`
	Published time.Time `json:"published" binding:"required"`
	ListPrice int       `json:"listPrice"`
}

// Books is a slice of Book structs
type Books map[string]Book
