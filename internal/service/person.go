package service

import (
	"github.com/Tonioou/go-api-test/internal/model"
	"github.com/Tonioou/go-api-test/internal/repository"
	"github.com/joomcode/errorx"
)

type Person interface {
	Post(person model.CreatePerson) (*model.Person, *errorx.Error)
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

func (ps *PersonService) Post(createPerson model.CreatePerson) (*model.Person, *errorx.Error) {
	person := model.NewPerson(createPerson.Age, createPerson.Email, createPerson.Name)
	err := ps.personRepository.Add(person)
	if err != nil {
		return &model.Person{}, err
	}
	return person, nil
}

func (ps *PersonService) Get() ([]model.Person, *errorx.Error) {
	return nil, nil
}
