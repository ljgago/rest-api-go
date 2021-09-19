package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"github.com/ljgago/api-rest/internal/core/domain"
	"github.com/ljgago/api-rest/internal/core/service"
	"github.com/ljgago/api-rest/internal/repository"
	"github.com/ljgago/api-rest/internal/routes"
)

func TestBookController(t *testing.T) {
	responseBooks := []domain.Book{
		{
			ID:     1,
			Title:  "The Lord of the Rings: The Fellowship of the Ring",
			Author: "J. R. R. Tolkien",
		},
		{
			ID:     2,
			Title:  "Harry Potter and the Philosopher's Stone",
			Author: "J. K. Rowling",
		},
	}

	responseBook := domain.Book{
		ID:     1,
		Title:  "The Lord of the Rings: The Fellowship of the Ring",
		Author: "J. R. R. Tolkien",
	}

	book := domain.Book{
		Title:  "The Lord of the Rings: The Fellowship of the Ring",
		Author: "J. R. R. Tolkien",
	}

	mockBookRepository := &repository.BookRepositoryMock{
		ReadFunc: func() ([]domain.Book, error) {
			return responseBooks, nil
		},
		ReadOneFunc: func(id int) (domain.Book, error) {
			return responseBook, nil
		},
		CreateFunc: func(book domain.Book) (domain.Book, error) {
			return responseBook, nil
		},
		UpdateFunc: func(id int, book domain.Book) (domain.Book, error) {
			return responseBook, nil
		},
		DeleteFunc: func(id int) error {
			return nil
		},
	}

	bookService := service.NewBookService(mockBookRepository)
	bookController := NewBookController(bookService)

	r := chi.NewRouter()
	routes.BookRoutes(r, bookController)

	t.Run("ListBooks", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
		res := httptest.NewRecorder()

		r.ServeHTTP(res, req)

		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.Code)

		data, err := json.Marshal(map[string][]domain.Book{"data": responseBooks})
		assert.NoError(t, err)

		// the json.NewEncoder(w).Encode(books) add a "\n" to end of json
		assert.Equal(t, string(data)+"\n", string(body))
	})

	t.Run("GetBook", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodGet, "/api/books/1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.Code)

		data, err := json.Marshal(map[string]domain.Book{"data": responseBook})
		assert.NoError(t, err)

		// the json.NewEncoder(w).Encode(books) add a "\n" to end of json
		assert.Equal(t, string(data)+"\n", string(body))
	})

	t.Run("CreateBook", func(t *testing.T) {
		t.Parallel()

		jsonBytes, err := json.Marshal(book)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPost, "/api/books", bytes.NewBuffer(jsonBytes))
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusCreated, res.Code)

		data, err := json.Marshal(map[string]domain.Book{"data": responseBook})
		assert.NoError(t, err)

		// the json.NewEncoder(w).Encode(books) add a "\n" to end of json
		assert.Equal(t, string(data)+"\n", string(body))
	})

	t.Run("UpdateBook", func(t *testing.T) {
		t.Parallel()

		jsonBytes, err := json.Marshal(book)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/api/books/1", bytes.NewBuffer(jsonBytes))
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, res.Code)

		data, err := json.Marshal(map[string]domain.Book{"data": responseBook})
		assert.NoError(t, err)

		// the json.NewEncoder(w).Encode(books) add a "\n" to end of json
		assert.Equal(t, string(data)+"\n", string(body))
	})

	t.Run("DeleteBook", func(t *testing.T) {
		t.Parallel()

		req := httptest.NewRequest(http.MethodDelete, "/api/books/1", nil)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)

		assert.Equal(t, http.StatusNoContent, res.Code)
	})
}
