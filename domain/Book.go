package domain

type Book struct {
	name string
}

func NewBook(name string) *Book {
	return &Book{name: name}
}
