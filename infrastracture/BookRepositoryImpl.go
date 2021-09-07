package infrastracture

import "github.com/inobu/practice-gin/domain"

type BookRepository struct {
	sqlConnection SqlConnection
}

func NewBookRepository(sqlConnection SqlConnection) domain.BookRepository {
	bookRepository := BookRepository{sqlConnection}
	return &bookRepository
}

func (bookRepository *BookRepository) GetBook(userId string) (domain.Book, error) {
	// TODO
}
