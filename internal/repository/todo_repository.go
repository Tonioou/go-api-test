package repository

import (
	"github.com/Tonioou/go-person-crud/internal/client"
	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Todo interface {
	GetById(id uuid.UUID) (model.Todo, *errorx.Error)
	Save(todo model.Todo) (model.Todo, *errorx.Error)
	Update(todo model.Todo) (model.Todo, *errorx.Error)
	Delete(id uuid.UUID) *errorx.Error
}

type TodoRepository struct {
	PgClient client.Postgres
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		PgClient: client.GetPgClient(),
	}
}

func (tr *TodoRepository) GetById(id uuid.UUID) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Save(todo model.Todo) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Update(todo model.Todo) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Delete(id uuid.UUID) *errorx.Error {
	return nil
}
