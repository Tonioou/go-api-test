package command

type DeleteTodo struct {
	ID string `param:"id" validate:"required,uuid"`
}
