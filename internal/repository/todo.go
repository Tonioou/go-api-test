package repository

import (
	"context"

	"github.com/Tonioou/go-person-crud/internal/client"
	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
	Save(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error)
	Delete(ctx context.Context, id uuid.UUID) *errorx.Error
}

type TodoRepository struct {
	PgClient client.Postgres
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{
		PgClient: client.GetPgClient(),
	}
}

func (tr *TodoRepository) GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Save(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Update(ctx context.Context, todo model.Todo) (model.Todo, *errorx.Error) {
	return model.Todo{}, nil
}

func (tr *TodoRepository) Delete(ctx context.Context, id uuid.UUID) *errorx.Error {
	return nil
}
