package service

import (
	"context"

	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, error)
	Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TodoRepository interface {
	TodoWriter
	TodoReader
}
type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService(repository TodoRepository) *TodoService {
	return &TodoService{
		todoRepository: repository,
	}
}

func (tr *TodoService) GetById(ctx context.Context, id uuid.UUID) (model.Todo, error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "GetById")
	defer span.End()
	return tr.todoRepository.GetById(newCtx, id)
}

func (tr *TodoService) Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Save")
	defer span.End()
	todo := model.NewTodo(addTodo.Name)
	return tr.todoRepository.Save(newCtx, todo)
}

func (tr *TodoService) Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Update")
	defer span.End()
	_, errx := tr.todoRepository.GetById(newCtx, updateTodo.Id)
	if errx != nil {
		span.AddEvent("Todo Id does not exist", trace.WithAttributes(attribute.String("todo-id", updateTodo.Id.String())))
		return model.Todo{}, errx
	}
	return tr.todoRepository.Update(newCtx, updateTodo)
}

func (tr *TodoService) Delete(ctx context.Context, id uuid.UUID) error {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Delete")
	defer span.End()
	_, errx := tr.todoRepository.GetById(newCtx, id)
	if errx != nil {
		span.AddEvent("Todo Id does not exist", trace.WithAttributes(attribute.String("todo-id", id.String())))
		return errx
	}
	return tr.todoRepository.Delete(newCtx, id)

}
