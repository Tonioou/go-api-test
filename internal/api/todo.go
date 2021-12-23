package api

import (
	"github.com/Tonioou/go-person-crud/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	v1.GET("/todo", ta.Get)
	v1.POST("/todo", ta.Save)
	v1.PUT("/todo/:id", ta.Update)
	v1.DELETE("/todo/:id", ta.Delete)
}

func (ta *TodoApi) Get(c *gin.Context) {
	value := c.Query("id")
	id, err := uuid.Parse(value)
	if err != nil {
		c.String(500, "invalid id")
	}
	result, errx := ta.TodoService.GetById(c.Request.Context(), id)
	if errx != nil {
		c.String(500, "invalid id")
	}
	c.JSON(200, result)
}

func (ta *TodoApi) Save(c *gin.Context) {
	c.String(201, "created")
}

func (ta *TodoApi) Update(c *gin.Context) {
	c.String(204, "updated")
}

func (ta *TodoApi) Delete(c *gin.Context) {
	c.String(204, "updated")
}
