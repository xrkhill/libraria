package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/xrkhill/libraria/internal/data"
)

func TestCreateBooks(t *testing.T) {
	testCases := []struct {
		requestMethod string
		requestURI    string
		requestBody   string
		wantCode      int
		wantBody      string
	}{
		{
			"POST",
			"/books",
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":999}`,
			201,
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":999}`,
		},
		{
			"POST",
			"/books",
			`{}`,
			400,
			`{"error":"Key: 'Book.Author' Error:Field validation for 'Author' failed on the 'required' tag\nKey: 'Book.Title' Error:Field validation for 'Title' failed on the 'required' tag\nKey: 'Book.ISBN' Error:Field validation for 'ISBN' failed on the 'required' tag\nKey: 'Book.Published' Error:Field validation for 'Published' failed on the 'required' tag"}`,
		},
	}

	defaultBooks := data.Books{}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test %s %s", testCase.requestMethod, testCase.requestURI), func(t *testing.T) {
			router := setupRouter(defaultBooks)
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(testCase.requestMethod, testCase.requestURI, strings.NewReader(testCase.requestBody))
			router.ServeHTTP(w, req)

			assert.Equal(t, testCase.wantCode, w.Code)
			assert.Equal(t, testCase.wantBody, w.Body.String())
		})
	}
}

func TestExistingBooks(t *testing.T) {
	testCases := []struct {
		requestMethod string
		requestURI    string
		requestBody   string
		wantCode      int
		wantBody      string
	}{
		{
			"POST",
			"/books",
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":999}`,
			304,
			``,
		},
		{
			"GET",
			"/books",
			``,
			200,
			`{"9780451524935":{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":999}}`,
		},
		{
			"GET",
			"/books/9780451524935",
			``,
			200,
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":999}`,
		},
		{
			"GET",
			"/books/4242",
			``,
			404,
			`{"error":"book 4242 does not exist, unable to read"}`,
		},
		{
			"PUT",
			"/books",
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":4242}`,
			200,
			`{"author":"George Orwell","title":"1984","ISBN":"9780451524935","language":"English","published":"1961-01-01T00:00:00Z","listPrice":4242}`,
		},
		{
			"DELETE",
			"/books/9780451524935",
			``,
			200,
			`{}`,
		},
	}

	defaultBooks := data.Books{
		"9780451524935": data.Book{
			Author:    "George Orwell",
			Title:     "1984",
			ISBN:      "9780451524935",
			Language:  "English",
			Published: time.Date(1961, time.January, 1, 0, 0, 0, 0, time.UTC),
			ListPrice: 999,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Test %s %s", testCase.requestMethod, testCase.requestURI), func(t *testing.T) {
			router := setupRouter(defaultBooks)
			w := httptest.NewRecorder()

			req, _ := http.NewRequest(testCase.requestMethod, testCase.requestURI, strings.NewReader(testCase.requestBody))
			router.ServeHTTP(w, req)

			assert.Equal(t, testCase.wantCode, w.Code)
			assert.Equal(t, testCase.wantBody, w.Body.String())
		})
	}
}
