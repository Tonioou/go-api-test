package service

import (
	"context"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type TodoWriter interface {
	Save(ctx context.Context, todo *model.Todo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error)
	Delete(ctx context.Context, id uuid.UUID) *errorx.Error
}
