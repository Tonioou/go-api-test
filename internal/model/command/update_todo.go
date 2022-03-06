package command

import (
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type UpdateTodo struct {
	Id   uuid.UUID `param:"id"`
	Name string    `json:"name"`
}

func (u *UpdateTodo) IsValid() *errorx.Error{
	if u.Id == uuid.Nil {
		return errorx.IllegalArgument.New("invalid id")
	}
	if u.Name == "" {
		return errorx.IllegalArgument.New("name cannot be empty")
	}
	return nil
}
