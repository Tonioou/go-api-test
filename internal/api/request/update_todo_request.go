package request

import (
	"github.com/Tonioou/go-person-crud/internal/model/command"
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
)

func InitializeUpdateTodo(c *gin.Context) (*command.UpdateTodo, *errorx.Error) {
	var updateTodo command.UpdateTodo
	if err := c.ShouldBindJSON(&updateTodo); err != nil {
		return &updateTodo, errorx.Decorate(err, "failed to bind data")
	}

	if err := updateTodo.IsValid(); err != nil {
		return &updateTodo, err
	}
	return &updateTodo, nil
}
