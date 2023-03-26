package service

import (
	"java-to-go/entity"
	"java-to-go/repository"
)

func NewPersonService(repo repository.PersonRepository) *PersonServiceImpl {
	return &PersonServiceImpl{repo: repo}
}

type PersonServiceImpl struct {
	repo repository.PersonRepository
}

func (p PersonServiceImpl) CreatePerson(person *entity.Person) (string, error) {
	return p.repo.CreatePerson(person)
}

func (p PersonServiceImpl) UpdatePerson(person *entity.Person, id string) (*entity.Person, error) {
	return p.repo.UpdatePerson(person, id)
}

func (p PersonServiceImpl) GetPersonById(id string) (*entity.Person, error) {
	return p.repo.GetPersonById(id)
}

func (p PersonServiceImpl) DeletePersonById(id string) error {
	return p.repo.DeletePersonById(id)
}
