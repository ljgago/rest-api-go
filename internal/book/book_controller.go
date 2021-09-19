package book

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// Controller
type Controller interface {
	ListBooks() http.HandlerFunc
	GetBook() http.HandlerFunc
	CreateBook() http.HandlerFunc
	UpdateBook() http.HandlerFunc
	DeleteBook() http.HandlerFunc
}

type controller struct {
	service Service
}

// NewBookController contructor for book controller.
func NewController(s Service) Controller {
	return &controller{service: s}
}

// ListBooks handler function for request a list all books from database.
// GET /api/books
func (c *controller) ListBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		books, err := c.service.Read()
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respondWithJSON(w, http.StatusOK, map[string][]Book{"data": books})
	}
}

// GetBook handler function for request to get a book from database.
// GET /api/books/{id}
func (c *controller) GetBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the book ID
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}
		book, err := c.service.ReadOne(int(id))
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusBadRequest, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]Book{"data": book})
	}
}

// CreateBook is a handler function for request add a book to database.
// POST /api/books <JSON Body>
func (c *controller) CreateBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Transform into RequestBody struct
		var book Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusInternalServerError, map[string]map[string]string{"errors": {"detail": err.Error()}})
			return
		}

		book, err = c.service.Create(book)
		if err != nil {
			log.Println(err)
			respondWithJSON(w, http.StatusBadRequest, map[string]map[string]string{"errors": {"detail": "Bad Request"}})
			return
		}

		respondWithJSON(w, http.StatusCreated, map[string]Book{"data": book})
	}
}

// UpdateBook is a handler function for request an update a book in database.
// PUT /api/books/{id} <JSON Body>
func (c *controller) UpdateBook() http.HandlerFunc {
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
		var book Book
		err = json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Update the book
		b, err := c.service.Update(id, book)
		if err != nil {
			log.Println("Hola", err)
			log.Fatalln(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respondWithJSON(w, http.StatusOK, map[string]Book{"data": b})
	}
}

// UpdateBook is a handler function for request to delete a book in database.
// DELETE /api/books/{id}
func (c *controller) DeleteBook() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Fatalln(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = c.service.Delete(id)
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
