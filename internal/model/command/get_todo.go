package command

type GetTodo struct {
	ID string `param:"id" validate:"required,uuid"`
}
