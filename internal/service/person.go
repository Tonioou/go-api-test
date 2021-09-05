package service

import (
	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/Tonioou/go-person-crud/internal/repository"
	"github.com/joomcode/errorx"
)

type Person interface {
	Post(person model.CreatePerson) (*model.Person, *errorx.Error)
	Get() ([]*model.Person, *errorx.Error)
	GetById(id string) (*model.Person, *errorx.Error)
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

func (ps *PersonService) Get() ([]*model.Person, *errorx.Error) {
	return ps.personRepository.Get()
}

func (ps *PersonService) GetById(id string) (*model.Person, *errorx.Error) {
	return ps.personRepository.GetById(id)
}
