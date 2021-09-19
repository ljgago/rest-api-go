package domain

// Book is the data structure that represents a book.
type Book struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

