package data

import (
	"time"
)

// Book is a struct for library book data
type Book struct {
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	ISBN      string    `json:"ISBN"`
	Language  string    `json:"language"`
	Published time.Time `json:"published"`
	ListPrice int       `json:"listPrice"`
}

// Books is a slice of Book structs
type Books map[string]Book
