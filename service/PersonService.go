package service

import (
	"java-to-go/entity"
	"java-to-go/repository"
)

type PersonService interface {
	CreatePerson(person *entity.Person) (string, error)
	UpdatePerson(person *entity.Person, id string) (*entity.Person, error)
	GetPersonById(id string) (*entity.Person, error)
	DeletePersonById(id string) (string, error)
}

type PersonServ struct {
	PersonService
}

func NewPersService(repo *repository.PersonRep) *PersonServ {
	return &PersonServ{
		PersonService: NewPersonService(repo.PersonRepository),
	}
}
