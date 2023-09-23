package request

import (
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func InitializeDeleteTodo(c echo.Context) (*command.DeleteTodo, error) {
	var deleteTodo command.DeleteTodo
	if err := c.Bind(&deleteTodo); err != nil {
		return &deleteTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := c.Validate(deleteTodo); err != nil {
		return &deleteTodo, err
	}
	return &deleteTodo, nil
}
