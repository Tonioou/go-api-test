package service

import (
	"github.com/Tonioou/go-api-test/model"
	"github.com/Tonioou/go-api-test/repository"
	"github.com/joomcode/errorx"
)

type Person interface {
	Post(person model.Person) (model.Person, *errorx.Error)
	Get() ([]model.Person, *errorx.Error)
}

type PersonService struct {
	personRepository repository.Person
}

func NewPersonService() *PersonService {
	return &PersonService{
		personRepository: repository.NewPersonRepository(),
	}
}

func (ps *PersonService) Post(person model.Person) (model.Person, *errorx.Error) {
	// return ps.personRepository.Add(person)
	return model.Person{}, nil
}

func (ps *PersonService) Get() ([]model.Person, *errorx.Error) {
	return nil, nil
}
