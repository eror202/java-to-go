package response

type PersonBookResponse struct {
	PersonId int64   `json:"personId"`
	BookId   []int64 `json:"bookId"`
}
