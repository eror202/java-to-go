package facade

import (
	"java-to-go/dto"
	"java-to-go/web/request"
	"java-to-go/web/response"
)

func CreateUserWithBooks(personBookRequest request.PersonBookRequest) *response.PersonBookResponse {
	PersonDto := dto.PersonDto{
		FullName: personBookRequest.PersonRequest.FullName,
		Title: personBookRequest.PersonRequest.Title,
		Age: personBookRequest.PersonRequest.Age,
	}
	
	return nil
}
