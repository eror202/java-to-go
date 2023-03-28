package request

type PersonRequest struct {
	FullName string `json:"fullName" db:"fullname"`
	Title    string `json:"title" db:"title"`
	Age      string `json:"age" db:"age"'`
}
