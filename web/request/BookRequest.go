package request

type BookRequest struct {
	Title     string `json:"title" db:"title"`
	Author    string `json:"author" db:"author"`
	PageCount string `json:"pageCount" db:"pageCount"`
}
