package dto

type BookDto struct {
	ID        string `json:"-" db:"id"`
	Title     string `json:"title" db:"title"`
	Author    string `json:"author" db:"author"`
	PageCount string `json:"pageCount" db:"pageCount"`
	PersonId  string `json:"personId" db:"personId"`
}
