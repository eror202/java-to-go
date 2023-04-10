package service

import (
	"java-to-go/databaseConfig"
	"java-to-go/dto"
	"java-to-go/entity"
	"java-to-go/exception"
	"java-to-go/repository"
)

func CreatePerson(personDto *dto.PersonDto) (string, error) {
	personToCreate := &entity.Person{
		FullName: personDto.FullName,
		Title:    personDto.Title,
		Age:      personDto.Age,
	}
	id, err := repository.NewPersRep(databaseConfig.ConnectToDb()).CreatePerson(personToCreate)
	if err != nil {
		return "", exception.NotCreatedObject{}
	}
	return id, nil
}

func UpdatePerson(personDto *dto.PersonDto, id string) (*dto.PersonDto, error) {
	personToUpdate := &entity.Person{
		FullName: personDto.FullName,
		Title:    personDto.Title,
		Age:      personDto.Age,
	}
	updatedPerson, err := repository.NewPersRep(databaseConfig.ConnectToDb()).UpdatePerson(personToUpdate, id)
	if err != nil {
		return nil, exception.NotFoundError{ID: id}
	}

	updatedPersonDto := &dto.PersonDto{
		FullName: updatedPerson.FullName,
		Title:    updatedPerson.Title,
		Age:      updatedPerson.Age,
	}
	return updatedPersonDto, nil
}

func GetPersonById(id string) (*dto.PersonDto, error) {
	person, err := repository.NewPersRep(databaseConfig.ConnectToDb()).GetPersonById(id)
	if err != nil {
		return nil, exception.NotFoundError{ID: id}
	}
	personDto := &dto.PersonDto{
		FullName: person.FullName,
		Title:    person.Title,
		Age:      person.Age,
	}
	return personDto, nil
}

func DeletePersonById(id string) (string, error) {

	_, err := repository.NewPersRep(databaseConfig.ConnectToDb()).DeletePersonById(id)
	if err != nil {
		return "", exception.NotFoundError{ID: id}
	}
	return "Объект удален", nil
}
