package service

import (
	"context"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/Tonioou/go-todo-list/internal/repository"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
	Save(ctx context.Context, addTodo *command.AddTodo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error)
	Delete(ctx context.Context, id uuid.UUID) *errorx.Error
}

type TodoRepository interface {
	TodoWriter
	TodoReader
}
type TodoService struct {
	todoRepository TodoRepository
}

func NewTodoService() *TodoService {
	return &TodoService{
		todoRepository: repository.NewTodoRepository(),
	}
}

func (tr *TodoService) GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "GetById")
	defer span.End()
	result, errx := tr.todoRepository.GetById(newCtx, id)
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
	result, errx := tr.todoRepository.Save(newCtx, todo)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
	}
	return result, errx
}

func (tr *TodoService) Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error) {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Update")
	defer span.End()
	_, errx := tr.todoRepository.GetById(newCtx, updateTodo.Id)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		span.AddEvent("Todo Id does not exist", trace.WithAttributes(attribute.String("todo-id", updateTodo.Id.String())))
		return model.Todo{}, errx
	}
	todo, errx := tr.todoRepository.Update(newCtx, updateTodo)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
	}
	return todo, errx
}

func (tr *TodoService) Delete(ctx context.Context, id uuid.UUID) *errorx.Error {
	newCtx, span := otel.Tracer("service-todo").Start(ctx, "Update")
	defer span.End()
	_, errx := tr.todoRepository.GetById(newCtx, id)
	if errx != nil {
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		span.AddEvent("Todo Id does not exist", trace.WithAttributes(attribute.String("todo-id", id.String())))
		return errx
	}
	return tr.todoRepository.Delete(newCtx, id)

}
