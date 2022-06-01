package service

import (
	"context"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type TodoReader interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
}
