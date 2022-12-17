package api

import (
	"github.com/Tonioou/go-todo-list/internal/api/request"
	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/service"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
)

type TodoApi struct {
	TodoService service.Todo
}

func NewTodoApi(todoService service.Todo) *TodoApi {
	return &TodoApi{
		TodoService: todoService,
	}
}

func (ta *TodoApi) Register(echo *echo.Echo) {
	v1 := echo.Group("/v1")

	v1.GET("/todo/:id", ta.GetById)
	v1.POST("/todo", ta.Save)
	v1.PUT("/todo/:id", ta.Update)
	v1.DELETE("/todo/:id", ta.Delete)
}

func (ta *TodoApi) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	value := c.Param("id")
	id, err := uuid.Parse(value)
	if err != nil {
		errorResponse := model.NewErrorResponse(errorx.Decorate(err, "failed to parse id"), config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}

	result, errx := ta.TodoService.GetById(ctx, id)

	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.JSON(200, result)
}

func (ta *TodoApi) Save(c echo.Context) error {
	ctx := c.Request().Context()

	addTodo, err := request.InitializeAddTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Save(ctx, addTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}

	return c.JSON(201, result)
}

func (ta *TodoApi) Update(c echo.Context) error {
	ctx := c.Request().Context()
	updateTodo, err := request.InitializeUpdateTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Update(ctx, updateTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.JSON(204, result)
}

func (ta *TodoApi) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	value := c.Param("id")
	id, err := uuid.Parse(value)
	if err != nil {
		errorResponse := model.NewErrorResponse(errorx.Decorate(err, "failed to parse id"), config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	errx := ta.TodoService.Delete(ctx, id)
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.String(204, "delete")
}
