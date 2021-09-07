package domain

type BookRepository interface {
	GetBook(userId string) (Book, error)
}
