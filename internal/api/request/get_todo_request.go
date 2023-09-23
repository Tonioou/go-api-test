package request

import (
	"github.com/Tonioou/go-todo-list/internal/model/command"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

func InitializeGetTodo(c echo.Context) (*command.GetTodo, error) {
	var getTodo command.GetTodo
	if err := c.Bind(&getTodo); err != nil {
		return &getTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := c.Validate(getTodo); err != nil {
		return &getTodo, err
	}
	return &getTodo, nil
}
