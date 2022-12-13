package api

import (
	"github.com/Tonioou/go-todo-list/internal/api/request"
	"github.com/Tonioou/go-todo-list/internal/config"
	"github.com/Tonioou/go-todo-list/internal/model"
	"github.com/Tonioou/go-todo-list/internal/service"
	"github.com/google/uuid"
	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
	}
	newCtx, span := otel.Tracer("api-todo").Start(ctx, "GetById", opts...)
	defer span.End()
	value := c.Param("id")
	id, err := uuid.Parse(value)
	if err != nil {
		errorResponse := model.NewErrorResponse(errorx.Decorate(err, "failed to parse id"), config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}

	result, errx := ta.TodoService.GetById(newCtx, id)

	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	//otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(c.Response().Header()))
	return c.JSON(200, result)
}

func (ta *TodoApi) Save(c echo.Context) error {
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
	}
	ctx := c.Request().Context()

	newCtx, span := otel.Tracer("api-todo").Start(ctx, "Save", opts...)
	defer span.End()
	addTodo, err := request.InitializeAddTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Save(newCtx, addTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)

	}
	span.AddEvent("what-is-an-event")

	return c.JSON(201, result)
}

func (ta *TodoApi) Update(c echo.Context) error {
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
	}
	ctx := c.Request().Context()
	newCtx, span := otel.Tracer("api-todo").Start(ctx, "Update", opts...)
	defer span.End()
	updateTodo, err := request.InitializeUpdateTodo(c)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	result, err := ta.TodoService.Update(newCtx, updateTodo)
	if err != nil {
		errorResponse := model.NewErrorResponse(err, config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.JSON(204, result)
}

func (ta *TodoApi) Delete(c echo.Context) error {
	opts := []trace.SpanStartOption{
		trace.WithSpanKind(trace.SpanKindClient),
	}
	ctx := c.Request().Context()
	newCtx, span := otel.Tracer("api-todo").Start(ctx, "Delete", opts...)
	defer span.End()
	value := c.Param("id")
	id, err := uuid.Parse(value)
	if err != nil {
		errorResponse := model.NewErrorResponse(errorx.Decorate(err, "failed to parse id"), config.Logger.Error)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	errx := ta.TodoService.Delete(newCtx, id)
	if errx != nil {
		errorResponse := model.NewErrorResponse(errx, config.Logger.Error)
		span.RecordError(errx)
		span.SetStatus(codes.Error, errx.Error())
		return c.JSON(errorResponse.StatusCode, errorResponse)
	}
	return c.String(204, "delete")
}
