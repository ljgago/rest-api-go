package port

import (
	"github.com/ljgago/api-rest/internal/core/domain"
)

// BookRespository
type BookRepository interface {
	Read() ([]domain.Book, error)
	ReadOne(id int) (domain.Book, error)
	Create(book domain.Book) (domain.Book, error)
	Update(id int, book domain.Book) (domain.Book, error)
	Delete(id int) error
}
