package service

import (
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/repository"
)

func NewPersonService(repo repository.PersonRepository) *PersonServiceImpl {
	return &PersonServiceImpl{repo: repo}
}

type PersonServiceImpl struct {
	repo repository.PersonRepository
}

func (p PersonServiceImpl) CreatePerson(personDto *dto.PersonDto) (string, error) {
	person := &entity.Person{
		FullName: personDto.FullName,
		Title:    personDto.Title,
		Age:      personDto.Age,
	}
	return p.repo.CreatePerson(person)
}

func (p PersonServiceImpl) UpdatePerson(personDto *dto.PersonDto, id string) (*dto.PersonDto, error) {
	return p.repo.UpdatePerson(person, id)
}

func (p PersonServiceImpl) GetPersonById(id string) (*dto.PersonDto, error) {
	return p.repo.GetPersonById(id)
}

func (p PersonServiceImpl) DeletePersonById(id string) (string, error) {
	return p.repo.DeletePersonById(id)
}
