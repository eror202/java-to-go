package repository

import (
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type BookRepPostgres struct {
	db *sqlx.DB
}

func (b BookRepPostgres) CreateBook(book *entity.Book) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookRepPostgres) UpdateBook(book *entity.Book, id string) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookRepPostgres) GetBookById(id string) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookRepPostgres) DeleteBookById(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepPostgres(db *sqlx.DB) *BookRepPostgres {
	return &BookRepPostgres{db: db}
}
