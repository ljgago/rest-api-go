package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/ljgago/api-rest/internal/core/domain"
	"github.com/ljgago/api-rest/internal/core/port"
)


type bookCtrl struct {
	bookServ port.BookService
}

// NewBookController contructor for book controller.
func NewBookController(b port.BookService) port.BookController {
	return &bookCtrl{bookServ: b}
}

// ListBooks handler function for request a list all books from database.
// GET /api/books
func (c *bookCtrl) ListBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := c.bookServ.Read()
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, http.StatusOK, map[string][]domain.Book{"data": books})
	}
}

// GetBook handler function for request to get a book from database.
// GET /api/books/{id}
func (c *bookCtrl) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the book ID
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}
		book, err := c.bookServ.ReadOne(int(id))
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusBadRequest, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]domain.Book{"data": book})
	}
}

// CreateBook is a handler function for request add a book to database.
// POST /api/books <JSON Body>
func (c *bookCtrl) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Transform into RequestBody struct
		var book domain.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}

		book, err = c.bookServ.Create(book)
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusBadRequest, map[string]map[string]string{"errors": {"detail": "Bad Request"}})
			return
		}

		respondWithJSON(w, http.StatusCreated, map[string]domain.Book{"data": book})
	}
}

// UpdateBook is a handler function for request an update a book in database.
// PUT /api/books/{id} <JSON Body>
func (c *bookCtrl) UpdateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// rawRequestBody, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// // Transform into RequestBody struct
		// var book Book
		// err = json.Unmarshal(rawRequestBody, &book)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	w.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// Read the book ID
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}
		// Transform into RequestBody struct
		var book domain.Book
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Update the book
		b, err := c.bookServ.Update(id, book)
		if err != nil {
			log.Println("Hola", err)
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]domain.Book{"data": b})
	}
}

// UpdateBook is a handler function for request to delete a book in database.
// DELETE /api/books/{id}
func (c *bookCtrl) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = c.bookServ.Delete(id)
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

