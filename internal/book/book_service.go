package book

type Service interface {
	Read() ([]Book, error)
	ReadOne(id int) (Book, error)
	Create(book Book) (Book, error)
	Update(id int, book Book) (Book, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

// NewService
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Read() ([]Book, error) {
	return s.repository.Read()
}

func (s *service) ReadOne(id int) (Book, error) {
	return s.repository.ReadOne(id)
}

func (s *service) Create(book Book) (Book, error) {
	return s.repository.Create(book)
}

func (s *service) Update(id int, book Book) (Book, error) {
	return s.repository.Update(id, book)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
