package service

import (
	"java-to-go/entity"
	"java-to-go/repository"
)

type PersonService interface {
	CreatePerson(person *entity.Person) (string, error)
	UpdatePerson(person *entity.Person, id string) (*entity.Person, error)
	GetPersonById(id string) (*entity.Person, error)
	DeletePersonById(id string) error
}

type Service struct {
	PersonService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		PersonService: NewPersonService(repo.PersonRepository),
	}
}
