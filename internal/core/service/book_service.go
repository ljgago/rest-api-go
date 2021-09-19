package service

import (
	"github.com/ljgago/api-rest/internal/core/domain"
	"github.com/ljgago/api-rest/internal/core/port"
)

type bookServ struct {
	bookRepo port.BookRepository
}

// NewBookService
func NewBookService(r port.BookRepository) port.BookService {
	return &bookServ{
		bookRepo: r,
	}
}

func (s *bookServ) Read() ([]domain.Book, error) {
	return s.bookRepo.Read()
}

func (s *bookServ) ReadOne(id int) (domain.Book, error) {
	return s.bookRepo.ReadOne(id)
}

func (s *bookServ) Create(book domain.Book) (domain.Book, error) {
	return s.bookRepo.Create(book)
}

func (s *bookServ) Update(id int, book domain.Book) (domain.Book, error) {
	return s.bookRepo.Update(id, book)
}

func (s *bookServ) Delete(id int) error {
	return s.bookRepo.Delete(id)
}
