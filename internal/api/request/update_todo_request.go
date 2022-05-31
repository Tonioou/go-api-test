package request

import (
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/joomcode/errorx"
)

func InitializeUpdateTodo(c echo.Context) (*command.UpdateTodo, *errorx.Error) {
	var updateTodo command.UpdateTodo
	if err := c.Bind(&updateTodo); err != nil {
		return &updateTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := updateTodo.IsValid(); err != nil {
		return &updateTodo, err
	}
	return &updateTodo, nil
}
