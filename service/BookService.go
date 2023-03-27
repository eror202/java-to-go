package service

import (
	"java-to-go/entity"
	"java-to-go/repository"
)

type BookService interface {
	CreateBook(book *entity.Book) (string, error)
	UpdateBook(book *entity.Book, id string) (*entity.Book, error)
	GetBookById(id string) (*entity.Book, error)
	DeleteBookById(id string) (string, error)
}

type BookServ struct {
	BookService
}

func NewBookServ(repo *repository.BookRep) *BookServ {
	return &BookServ{
		BookService: NewBookService(repo.BookRepository),
	}
}
