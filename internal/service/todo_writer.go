package service

import (
	"context"

	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/google/uuid"
)

type TodoWriter interface {
	Save(ctx context.Context, todo *model.Todo) (model.Todo, error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
