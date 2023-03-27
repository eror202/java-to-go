package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type PersonRepPostgres struct {
	db *sqlx.DB
}

func (r *PersonRepPostgres) CreatePerson(person *entity.Person) (string, error) {
	var id string
	query := fmt.Sprintf("INSERT INTO %s (fullName, title, age) VALUES ($1, $2, $3) RETURNING ID", "Person")
	row := r.db.QueryRow(query, person.FullName, person.Title, person.Age)
	if err := row.Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *PersonRepPostgres) UpdatePerson(person *entity.Person, id string) (*entity.Person, error) {
	var updatePerson entity.Person
	query := fmt.Sprintf("UPDATE Person SET fullName = '%s', title = '%s', age = '%s' WHERE id = '%s'", person.FullName,
		person.Title, person.Age, id)
	if err := r.db.Get(&person, query); err != nil {
		return nil, err
	}
	return &updatePerson, nil
}

func (r *PersonRepPostgres) GetPersonById(id string) (*entity.Person, error) {
	var person entity.Person
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = '%s'", "Person", id)
	if err := r.db.Get(&person, query); err != nil {
		return nil, err
	}
	return &person, nil
}

func (r *PersonRepPostgres) DeletePersonById(id string) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = '%s'", "Person", id)
	if err := r.db.Query(query); err != nil {

	}
	return "", err
}

func NewPersonRepPostgres(db *sqlx.DB) *PersonRepPostgres {
	return &PersonRepPostgres{db: db}
}
