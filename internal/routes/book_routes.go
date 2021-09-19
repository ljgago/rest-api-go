package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/ljgago/api-rest/internal/core/port"
)

func BookRoutes(r *chi.Mux, bookController port.BookController) {
	r.Get("/api/books", bookController.ListBooks())
	r.Get("/api/books/{id}", bookController.GetBook())
	r.Post("/api/books", bookController.CreateBook())
	r.Put("/api/books/{id}", bookController.UpdateBook())
	r.Delete("/api/books/{id}", bookController.DeleteBook())
}
