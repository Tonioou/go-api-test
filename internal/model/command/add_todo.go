package command

import "github.com/joomcode/errorx"

type AddTodo struct {
	Name string `json:"name" validate:"required"`
}

func (at AddTodo) IsValid() *errorx.Error {
	if at.Name == "" {
		return errorx.IllegalArgument.New("invalid name")
	}
	return nil
}
