package entity

type Person struct {
	ID       string `json:"-" db:"id"`
	FullName string `json:"fullName" db:"full_name"`
	Title    string `json:"title" db:"title"`
	Age      string `json:"age" db:"age"`
}
