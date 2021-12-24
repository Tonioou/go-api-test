package api

import (
	"github.com/Tonioou/go-person-crud/internal/api/request"
	"github.com/Tonioou/go-person-crud/internal/config"
	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/Tonioou/go-person-crud/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
)

type TodoApi struct {
	TodoService service.Todo
}

func NewTodoApi() *TodoApi {
	return &TodoApi{
		TodoService: service.NewTodoService(),
	}
}

func (ta *TodoApi) Register(gin *gin.Engine) {
	v1 := gin.Group("/v1")

	v1.GET("/todo/:id", ta.GetById)
	v1.POST("/todo", ta.Save)
	v1.PUT("/todo/:id", ta.Update)
	v1.DELETE("/todo/:id", ta.Delete)
}

func (ta *TodoApi) GetById(c *gin.Context) {
	ctx := c.Request.Context()
	value := c.Param("id")
	id, err := uuid.Parse(value)
	if err != nil {
		errorResponse := model.NewErrorResponse(errorx.Decorate(err, "failed to parse id"), config.Logger.Warn)
		c.JSON(errorResponse.StatusCode, errorResponse)
		return
	}
	result, errx := ta.TodoService.GetById(ctx, id)
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		c.JSON(errorResponse.StatusCode, errorResponse)
		return
	}
	c.JSON(200, result)
}

func (ta *TodoApi) Save(c *gin.Context) {
	ctx := c.Request.Context()
	addTodo, errx := request.InitializeAddTodo(c)
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Warn)
		c.JSON(errorResponse.StatusCode, errorResponse)
		return
	}
	result, errx := ta.TodoService.Save(ctx, addTodo)
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		c.JSON(errorResponse.StatusCode, errorResponse)
		return
	}
	c.JSON(201, result)
}

func (ta *TodoApi) Update(c *gin.Context) {
	c.String(204, "updated")
}

func (ta *TodoApi) Delete(c *gin.Context) {
	c.String(204, "updated")
}
