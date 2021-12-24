package service

import (
	"context"

	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/Tonioou/go-person-crud/internal/model/command"
	"github.com/Tonioou/go-person-crud/internal/repository"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
	Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error)
	Delete(ctx context.Context, id uuid.UUID) *errorx.Error
}

type TodoService struct {
	TodoRepository repository.Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		TodoRepository: repository.NewTodoRepository(),
	}
}

func (tr *TodoService) GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error) {
	result, errx := tr.TodoRepository.GetById(ctx, id)
	return result, errx
}

func (tr *TodoService) Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, *errorx.Error) {
	todo := model.NewTodo(addTodo.Name)
	result, errx := tr.TodoRepository.Save(ctx, todo)
	return result, errx
}

func (tr *TodoService) Update(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoService) Delete(ctx context.Context, id uuid.UUID) *errorx.Error {
	return nil
}
