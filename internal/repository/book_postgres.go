package repository

import (
	"database/sql"

	"github.com/ljgago/api-rest/internal/core/domain"
	"github.com/ljgago/api-rest/internal/core/port"
)

type bookRepo struct {
	db *sql.DB
}

// NewBookRepository
func NewBookRepository(db *sql.DB) port.BookRepository {
	return &bookRepo{db: db}
}

// Read read all books from database.
func (r *bookRepo) Read() ([]domain.Book, error) {
	query := "SELECT * FROM book"
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.Book{}, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return []domain.Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

// ReadOne read a book from database by id.
func (r *bookRepo) ReadOne(id int) (domain.Book, error) {
	query := "SELECT * FROM book WHERE id = $1"
	rows := r.db.QueryRow(query, id)

	var b domain.Book
	err := rows.Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		return domain.Book{}, err
	}

	return b, nil

}

// Create add a new book in the database.
func (r *bookRepo) Create(book domain.Book) (domain.Book, error) {
	query := "INSERT INTO book (title, author) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, book.Title, book.Author).Scan(&book.ID)
	if err != nil {
		return domain.Book{}, err
	}
	return book, nil
}

// Update update book in the database.
func (r *bookRepo) Update(id int, book domain.Book) (domain.Book, error) {
	query := `UPDATE book SET title = $2, author = $3 WHERE id = $1 RETURNING id, title, author`
	var b domain.Book
	err := r.db.QueryRow(query, id, book.Title, book.Author).Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		return domain.Book{}, err
	}
	return b, err
}

// Delete remove a book from database by id.
func (r *bookRepo) Delete(id int) error {
	query := `
	DELETE FROM book
	WHERE id=$1;
	`
	_, err := r.db.Exec(query, id)

	return err
}

