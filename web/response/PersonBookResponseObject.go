package response

import "java-to-go/dto"

type PersonBookResponseObject struct {
	Person dto.PersonDto `json:"personId"`
	Book   []dto.BookDto `json:"bookId"`
}
