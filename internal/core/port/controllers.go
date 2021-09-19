package port

import (
	"net/http"
)

// BookController
type BookController interface {
	ListBooks() http.HandlerFunc
	GetBook() http.HandlerFunc
	CreateBook() http.HandlerFunc
	UpdateBook() http.HandlerFunc
	DeleteBook() http.HandlerFunc
}
