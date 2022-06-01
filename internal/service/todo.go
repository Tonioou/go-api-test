package service

import (
	"context"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/Tonioou/go-todo-list/internal/repository"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
	Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error)
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
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "GetById")
	defer span.End()
	result, errx := tr.TodoRepository.GetById(newCtx, id)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
	}
	return result, errx
}

func (tr *TodoService) Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, *errorx.Error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Save")
	defer span.End()
	todo := model.NewTodo(addTodo.Name)
	result, errx := tr.TodoRepository.Save(newCtx, todo)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
	}
	return result, errx
}

func (tr *TodoService) Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Update")
	defer span.End()
	_, errx := tr.TodoRepository.GetById(newCtx, updateTodo.Id)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		return model.Todo{}, errx
	}
	todo, errx := tr.TodoRepository.Update(newCtx, updateTodo)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
	}
	return todo, errx
}

func (tr *TodoService) Delete(ctx context.Context, id uuid.UUID) *errorx.Error {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Update")
	defer span.End()
	_, errx := tr.TodoRepository.GetById(newCtx, id)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		return errx
	}
	return tr.TodoRepository.Delete(newCtx, id)

}
