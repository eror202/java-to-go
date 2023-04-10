package service

import (
	"java-to-go/databaseConfig"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/exception"
	"java-to-go/repository"
)

func CreateBook(book *dto.BookDto) (string, error) {
	bookToCreate := &entity.Book{
		Title:     book.Title,
		Author:    book.Author,
		PageCount: book.PageCount,
		PersonId:  book.PersonId,
	}
	id, err := repository.NewBookRep(databaseConfig.ConnectToDb()).CreateBook(bookToCreate)
	if err != nil {
		return "", exception.NotCreatedObject{}
	}
	return id, nil
}

func UpdateBook(book *dto.BookDto, id string) (*dto.BookDto, error) {
	bookToUpdate := &entity.Book{
		Title:     book.Title,
		Author:    book.Author,
		PageCount: book.PageCount,
		PersonId:  book.PersonId,
	}
	updatedBook, err := repository.NewBookRep(databaseConfig.ConnectToDb()).UpdateBook(bookToUpdate, id)
	if err != nil {
		return nil, exception.NotFoundError{ID: id}
	}
	updatedBookDto := &dto.BookDto{
		Title:     updatedBook.Title,
		Author:    updatedBook.Author,
		PageCount: updatedBook.PageCount,
		PersonId:  updatedBook.PersonId,
	}
	return updatedBookDto, nil
}

func GetBookById(id string) (*dto.BookDto, error) {
	book, err := repository.NewBookRep(databaseConfig.ConnectToDb()).GetBookById(id)
	if err != nil {
		return nil, exception.NotFoundError{ID: id}
	}
	bookDto := &dto.BookDto{
		Title:     book.Title,
		Author:    book.Author,
		PageCount: book.PageCount,
		PersonId:  book.PersonId,
	}
	return bookDto, nil
}

func DeleteBookById(id string) (string, error) {
	_, err := repository.NewBookRep(databaseConfig.ConnectToDb()).DeleteBookById(id)
	if err != nil {
		return "", exception.NotFoundError{ID: id}
	}
	return "", nil
}

func DeleteBookByPersonId(id string) (string, error) {
	_, err := repository.NewBookRep(databaseConfig.ConnectToDb()).DeleteBookByPersonId(id)
	if err != nil {
		return "", exception.NotFoundError{ID: id}
	}
	return "", nil
}

func GetBookByPersonId(id string) ([]dto.BookDto, error) {
	var bookIdDto []dto.BookDto
	BookId, err := repository.NewBookRep(databaseConfig.ConnectToDb()).GetBookByPersonId(id)
	if err != nil {
		return nil, exception.NotFoundError{ID: id}
	}
	for _, value := range *BookId {
		bo := &dto.BookDto{
			Title:     value.Title,
			Author:    value.Author,
			PageCount: value.PageCount,
			PersonId:  value.PersonId,
		}
		bookIdDto = append(bookIdDto, *bo)
	}

	return bookIdDto, nil
}
