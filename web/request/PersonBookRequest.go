package request

type PersonBookRequest struct {
	PersonRequest PersonRequest `json:"personRequest"`
	BookRequest   []BookRequest `json:"bookRequest"`
}
