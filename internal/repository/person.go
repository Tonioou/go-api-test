package repository

import (
	"github.com/Tonioou/go-api-test/internal/dao"
	"github.com/Tonioou/go-api-test/internal/model"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Person interface {
	Add(person *model.Person) *errorx.Error
	GetById(id uuid.UUID) (model.Person, *errorx.Error)
	Get() ([]model.Person, *errorx.Error)
}

type PersonRepository struct {
	database  dao.Database
	tableName string
}

func NewPersonRepository() *PersonRepository {
	return &PersonRepository{
		database:  dao.GetDatabaseInMemoryDatabase(),
		tableName: "person",
	}
}

func (pr *PersonRepository) Add(person *model.Person) *errorx.Error {
	return pr.database.Add(pr.tableName, interface{}(person))
}

func (pr *PersonRepository) GetById(id uuid.UUID) (model.Person, *errorx.Error) {
	return model.Person{}, nil
}

func (pr *PersonRepository) Get() ([]model.Person, *errorx.Error) {
	return nil, nil
}
