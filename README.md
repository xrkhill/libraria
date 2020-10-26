# Libraria

Libraria is a sample REST API for books.

## API

Following are some example API requests and responses.

### GET /books

```bash
curl -s -X GET http://localhost:8080/books
```

```json
{
  "9780451524935": {
    "author": "George Orwell",
    "title": "1984",
    "ISBN": "978-0451524935",
    "language": "English",
    "published": "1961-01-01T00:00:00Z",
    "listPrice": 999
  },
  "9780553380163": {
    "author": "Stephen Hawking",
    "title": "A Brief History of Time",
    "ISBN": "978-0553380163",
    "language": "English",
    "published": "1998-09-01T00:00:00Z",
    "listPrice": 1800
  }
}
```

### POST /books

```bash
curl -s -X POST http://localhost:8080/books -H "Content-Type: application/json" -d "@sample_create.json"
```
