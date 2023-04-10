package entity

type Book struct {
	ID        string `json:"-" db:"id"`
	Title     string `json:"title" db:"title"`
	Author    string `json:"author" db:"author"`
	PageCount string `json:"pageCount" db:"page_count"`
	PersonId  string `json:"personId" db:"person_id"`
}
