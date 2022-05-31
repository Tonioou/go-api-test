package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Tonioou/go-todo-list/internal/model/command"

	"github.com/Tonioou/go-todo-list/internal/client"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/joomcode/errorx"
)

type Todo interface {
	GetById(ctx context.Context, id uuid.UUID) (model.Todo, *errorx.Error)
	Save(ctx context.Context, todo *model.Todo) (model.Todo, *errorx.Error)
	Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error)
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
	result := model.Todo{}
	query := `SELECT id, 
					name,
					created_at,
					finished,
					finished_at,
					active
				FROM todo
				WHERE id=$1;`

	row, errx := tr.PgClient.QueryRow(ctx, query, id)
	if errx != nil {
		return model.Todo{}, errx
	}
	var sqlTime sql.NullTime
	args := []interface{}{
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.Finished,
		&sqlTime,
		&result.Active,
	}
	err := row.Scan(args...)
	result.FinishedAt = sqlTime.Time
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return model.Todo{}, model.NotFound.New("todo not found")
		}
		return model.Todo{}, errorx.Decorate(err, "failed to scan row")
	}
	return result, nil
}

func (tr *TodoRepository) Save(ctx context.Context, todo *model.Todo) (model.Todo, *errorx.Error) {
	query := "INSERT INTO todo (id, name, created_at, finished, active) VALUES ($1,$2,$3,$4,$5);"

	id := uuid.New()
	args := []interface{}{
		&id,
		&todo.Name,
		&todo.CreatedAt,
		&todo.Finished,
		&todo.Active,
	}

	errx := tr.PgClient.Exec(ctx, query, args...)
	if errx != nil {
		return model.Todo{}, errx
	}
	return tr.GetById(ctx, id)
}

func (tr *TodoRepository) Update(ctx context.Context, updateTodo *command.UpdateTodo) (model.Todo, *errorx.Error) {
	query := "UPDATE todo SET name=$1 where id=$2;"

	args := []interface{}{
		&updateTodo.Name,
		&updateTodo.Id,
	}

	errx := tr.PgClient.Exec(ctx, query, args...)
	if errx != nil {
		return model.Todo{}, errx
	}
	return tr.GetById(ctx, updateTodo.Id)
}

func (tr *TodoRepository) Delete(ctx context.Context, id uuid.UUID) *errorx.Error {
	query := "DELETE FROM  todo  where id=$1;"
	args := []interface{}{
		&id,
	}

	errx := tr.PgClient.Exec(ctx, query, args...)
	return errx
}
