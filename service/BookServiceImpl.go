package service

import (
	"java-to-go/entity"
	"java-to-go/repository"
)

func NewBookService(repo repository.BookRepository) *BookServiceImpl {
	return &BookServiceImpl{repo: repo}
}

type BookServiceImpl struct {
	repo repository.BookRepository
}

func (b BookServiceImpl) CreateBook(book *entity.Book) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookServiceImpl) UpdateBook(book *entity.Book, id string) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookServiceImpl) GetBookById(id string) (*entity.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b BookServiceImpl) DeleteBookById(id string) (string, error) {
	//TODO implement me
	panic("implement me")
}
