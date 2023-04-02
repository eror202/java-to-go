package response

type PersonBookResponse struct {
	PersonId string   `json:"personId"`
	BookId   []string `json:"bookId"`
}
