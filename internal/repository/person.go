package repository

import (
	"github.com/Tonioou/go-api-test/internal/dao"
	"github.com/Tonioou/go-api-test/internal/model"
	"github.com/joomcode/errorx"
)

type Person interface {
	Add(person *model.Person) *errorx.Error
	GetById(id string) (*model.Person, *errorx.Error)
	Get() ([]*model.Person, *errorx.Error)
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

func (pr *PersonRepository) GetById(id string) (*model.Person, *errorx.Error) {
	item, err := pr.database.GetById(pr.tableName, id)
	if err != nil {
		return nil, err
	}
	if item == nil {
		return &model.Person{}, nil
	}
	return item.(*model.Person), nil
}

func (pr *PersonRepository) Get() ([]*model.Person, *errorx.Error) {
	people := make([]*model.Person, 0)
	items, err := pr.database.Get(pr.tableName)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		person := item.(*model.Person)
		people = append(people, person)
	}
	return people, nil
}
