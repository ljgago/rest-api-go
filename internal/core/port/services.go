package port

import (
	"github.com/ljgago/api-rest/internal/core/domain"
)

// BookService
type BookService interface {
	Read() ([]domain.Book, error)
	ReadOne(id int) (domain.Book, error)
	Create(book domain.Book) (domain.Book, error)
	Update(id int, book domain.Book) (domain.Book, error)
	Delete(id int) error
}

