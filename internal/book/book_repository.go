package book

import (
	"database/sql"
)

type Repository interface {
	Read() ([]Book, error)
	ReadOne(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(id int, book Book) (Book, error)
	Delete(id int) error
}

// Data Sources

type repository struct {
	db *sql.DB
}

// NewRepository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

// Read read all books from database.
func (r *repository) Read() ([]Book, error) {
	query := "SELECT * FROM book"
	rows, err := r.db.Query(query)
	if err != nil {
		return []Book{}, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author)
		if err != nil {
			return []Book{}, err
		}

		books = append(books, book)
	}

	return books, nil
}

// ReadOne read a book from database by id.
func (r *repository) ReadOne(id int) (Book, error) {
	query := "SELECT * FROM book WHERE id = $1"
	rows := r.db.QueryRow(query, id)

	var b Book
	err := rows.Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		return Book{}, err
	}

	return b, nil

}

// Create add a new book in the database.
func (r *repository) Create(book Book) (Book, error) {
	query := "INSERT INTO book (title, author) VALUES ($1, $2) RETURNING id"
	err := r.db.QueryRow(query, book.Title, book.Author).Scan(&book.ID)
	if err != nil {
		return Book{}, err
	}
	return book, nil
}

// Update update book in the database.
func (r *repository) Update(id int, book Book) (Book, error) {
	query := `UPDATE book SET title = $2, author = $3 WHERE id = $1 RETURNING id, title, author`
	var b Book
	err := r.db.QueryRow(query, id, book.Title, book.Author).Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		return Book{}, err
	}
	return b, err
}

// Delete remove a book from database by id.
func (r *repository) Delete(id int) error {
	query := `
	DELETE FROM book
	WHERE id=$1;
	`
	_, err := r.db.Exec(query, id)

	return err
}
