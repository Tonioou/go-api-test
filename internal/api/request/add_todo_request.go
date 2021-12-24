package request

import (
	"github.com/Tonioou/go-person-crud/internal/model/command"
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
)

func InitializeAddTodo(c *gin.Context) (*command.AddTodo, *errorx.Error) {
	var addTodo command.AddTodo
	if err := c.ShouldBindJSON(&addTodo); err != nil {
		return &addTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := addTodo.IsValid(); err != nil {
		return &addTodo, err
	}
	return &addTodo, nil
}
