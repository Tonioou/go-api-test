package api

import (
	"github.com/Tonioou/go-todo-list/internal/api/request"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TodoApi struct {
	TodoService service.Todo
}

func NewTodoApi(todoService service.Todo) *TodoApi {
	return &TodoApi{
		TodoService: todoService,
	}
}

func (ta *TodoApi) Register(e *echo.Echo) {
	v1 := e.Group("/v1")

	v1.GET("/todo/:id", ta.GetById)
	v1.POST("/todo", ta.Save)
	v1.PUT("/todo/:id", ta.Update)
	v1.DELETE("/todo/:id", ta.Delete)
}

func (ta *TodoApi) GetById(c echo.Context) error {
	ctx := c.Request().Context()

	getTodo, err := request.InitializeGetTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}

	result, errx := ta.TodoService.GetById(ctx, uuid.MustParse(getTodo.ID))

	if errx != nil {
		errorResponse := model.NewErrorResponse(errx)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.JSON(http.StatusOK, result)
}

func (ta *TodoApi) Save(c echo.Context) error {
	ctx := c.Request().Context()

	addTodo, err := request.InitializeAddTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Save(ctx, addTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}

	return c.JSON(http.StatusCreated, result)
}

func (ta *TodoApi) Update(c echo.Context) error {
	ctx := c.Request().Context()
	updateTodo, err := request.InitializeUpdateTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Update(ctx, updateTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.JSON(http.StatusOK, result)
}

func (ta *TodoApi) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	deleteTodo, err := request.InitializeDeleteTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	errx := ta.TodoService.Delete(ctx, uuid.MustParse(deleteTodo.ID))
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx)
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.NoContent(http.StatusNoContent)
}
