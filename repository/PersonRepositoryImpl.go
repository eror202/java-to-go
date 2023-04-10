package repository

import (
	"github.com/jmoiron/sqlx"
	"java-to-go/entity"
)

type PersonRepPostgres struct {
	db *sqlx.DB
}

func (r *PersonRepPostgres) CreatePerson(person *entity.Person) (string, error) {
	var id string
	if err := r.db.QueryRow("INSERT INTO Person (full_name, title, age) VALUES ($1, $2, $3) RETURNING ID", person.FullName, person.Title, person.Age).Scan(&id); err != nil {
		return "", err
	}
	return id, nil
}

func (r *PersonRepPostgres) UpdatePerson(person *entity.Person, id string) (*entity.Person, error) {

	if _, err := r.db.Query("UPDATE Person SET full_name = $1, title = $2, age = $3 WHERE id = $4", person.FullName, person.Title, person.Age, id); err != nil {
		return nil, err
	}

	return person, nil
}

func (r *PersonRepPostgres) GetPersonById(id string) (*entity.Person, error) {
	var person entity.Person
	if err := r.db.QueryRow("SELECT * FROM Person WHERE id = $1", id).Scan(&person.ID, &person.FullName, &person.Title, &person.Age); err != nil {
		return nil, err
	}

	return &person, nil
}

func (r *PersonRepPostgres) DeletePersonById(id string) (string, error) {
	if _, err := r.db.Exec("DELETE FROM Person WHERE id = $1", id); err != nil {
		return "", err
	}
	return "", nil
}

func NewPersonRepPostgres(db *sqlx.DB) *PersonRepPostgres {
	return &PersonRepPostgres{db: db}
}
