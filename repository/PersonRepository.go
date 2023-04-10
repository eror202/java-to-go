package repository

import (
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type PersonRepository interface {
	CreatePerson(person *entity.Person) (string, error)
	UpdatePerson(person *entity.Person, id string) (*entity.Person, error)
	GetPersonById(id string) (*entity.Person, error)
	DeletePersonById(id string) (string, error)
}

type PersonRep struct {
	PersonRepository
}

func NewPersRep(db *sqlx.DB) *PersonRep {
	return &PersonRep{
		PersonRepository: NewPersonRepPostgres(db),
	}
}
