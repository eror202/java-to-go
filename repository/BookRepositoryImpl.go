package repository

import (
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type BookRepPostgres struct {
	db *sqlx.DB
}

func (b BookRepPostgres) CreateBook(book *entity.Book) (string, error) {
	var id string
	if err := b.db.QueryRow("INSERT INTO Book (title,author, page_count, person_id) VALUES ($1, $2, $3, $4) RETURNING ID", book.Title, book.Author, book.PageCount,
		book.PersonId).Scan(&id); err != nil {
		return "", err
	}

	return id, nil
}

func (b BookRepPostgres) UpdateBook(book *entity.Book, id string) (*entity.Book, error) {
	var updatedBook entity.Book
	if err := b.db.QueryRow("UPDATE Book SET title = $1, author = $2, page_count = $3, person_id = $4",
		book.Title, book.Author, book.PageCount, book.PersonId).Scan(&updatedBook); err != nil {
		return nil, err
	}
	return &updatedBook, nil
}

func (b BookRepPostgres) GetBookById(id string) (*entity.Book, error) {
	var book entity.Book
	if err := b.db.QueryRow("SELECT * FROM Book WHERE id = $1", id).Scan(&book.ID, &book.Title, &book.Author, &book.PageCount, &book.PersonId); err != nil {
		return nil, err
	}

	return &book, nil
}

func (b BookRepPostgres) DeleteBookById(id string) (string, error) {
	if _, err := b.db.Exec("DELETE FROM Book WHERE id = $1", id); err != nil {
		return "", err
	}
	return "", nil
}

func (b BookRepPostgres) GetBookByPersonId(id string) (*[]entity.Book, error) {
	var bookId []entity.Book
	rows, err := b.db.Query("SELECT * FROM Book WHERE person_id = $1", id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		book := entity.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PageCount, &book.PersonId); err != nil {
			return nil, err
		}

		bookId = append(bookId, book)
	}
	return &bookId, nil
}

func (b BookRepPostgres) DeleteBookByPersonId(id string) (string, error) {
	if _, err := b.db.Exec("DELETE FROM Book WHERE person_id = $1", id); err != nil {
		return "", err
	}
	return "", nil
}

func NewBookRepPostgres(db *sqlx.DB) *BookRepPostgres {
	return &BookRepPostgres{db: db}
}
