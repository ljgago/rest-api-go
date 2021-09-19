package book

import (
	"github.com/go-chi/chi/v5"
)

func Routes(r *chi.Mux, controller Controller) {
	r.Get("/api/books", controller.ListBooks())
	r.Get("/api/books/{id}", controller.GetBook())
	r.Post("/api/books", controller.CreateBook())
	r.Put("/api/books/{id}", controller.UpdateBook())
	r.Delete("/api/books/{id}", controller.DeleteBook())
}
