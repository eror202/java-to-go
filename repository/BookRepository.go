package repository

import (
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type BookRepository interface {
	CreateBook(book *entity.Book) (string, error)
	UpdateBook(book *entity.Book, id string) (*entity.Book, error)
	GetBookById(id string) (*entity.Book, error)
	DeleteBookById(id string) (string, error)
	GetBookByPersonId(id string) (*[]entity.Book, error)
	DeleteBookByPersonId(id string) (string, error)
}

type BookRep struct {
	BookRepository
}

func NewBookRep(db *sqlx.DB) *BookRep {
	return &BookRep{BookRepository: NewBookRepPostgres(db)}
}
