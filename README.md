# Libraria

Libraria is a sample REST API for books.

## Running

Tested with Go 1.13.x. Uses Go Modules.

```bash
go test ./...
go run main.go
```

## API

Following are some example API requests and responses.

### POST /books

Create a book.

#### Request

```bash
curl -s -X POST http://localhost:8080/books -H "Content-Type: application/json" -d "@book_00.json"
```

#### Response

```json
{
  "author": "George Orwell",
  "title": "1984",
  "ISBN": "9780451524935",
  "language": "English",
  "published": "1961-01-01T00:00:00Z",
  "listPrice": 999
}
```

### GET /books

Read all books.

#### Request

```bash
curl -s -X GET http://localhost:8080/books
```

#### Response

```json
{
  "9780451524935": {
    "author": "George Orwell",
    "title": "1984",
    "ISBN": "9780451524935",
    "language": "English",
    "published": "1961-01-01T00:00:00Z",
    "listPrice": 999
  },
  "9780553380163": {
    "author": "Stephen Hawking",
    "title": "A Brief History of Time",
    "ISBN": "9780553380163",
    "language": "English",
    "published": "1998-09-01T00:00:00Z",
    "listPrice": 1800
  }
}
```

### GET /books/:isbn

Read one book.

#### Request

```bash
curl -s -X GET http://localhost:8080/books/9780451524935
```

#### Response

```json
{
  "author": "George Orwell",
  "title": "1984",
  "ISBN": "9780451524935",
  "language": "English",
  "published": "1961-01-01T00:00:00Z",
  "listPrice": 999
}
```

### PUT /books

Update a book.

#### Request

```bash
curl -s -X PUT http://localhost:8080/books -H "Content-Type: application/json" -d "@book_01.json"
```

#### Response

```json
{
  "author": "Stephen Hawking",
  "title": "A Brief History of Time",
  "ISBN": "9780553380163",
  "language": "English",
  "published": "1998-09-01T00:00:00Z",
  "listPrice": 1800
}
```

### DELETE /books/:isbn

Delete a book.

#### Request

```bash
curl -s -X DELETE http://localhost:8080/books/9780451524935
```

#### Response

```json
{}
```
