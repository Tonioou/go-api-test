package api

import (
	"net/http"

	"github.com/Tonioou/go-person-crud/internal/model"
	"github.com/Tonioou/go-person-crud/internal/service"
	"github.com/gin-gonic/gin"
)

type PersonApi struct {
	service service.Person
}

func NewPersonApi() *PersonApi {
	return &PersonApi{
		service.NewPersonService(),
	}
}

func (p *PersonApi) Register(e *gin.Engine) {
	person := e.Group("/person")
	person.GET("/", p.Get)
	person.GET("/:id", p.GetById)
	person.POST("/", p.Post)
}

func (p *PersonApi) Get(c *gin.Context) {
	people, err := p.service.Get()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Cause().Error()})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (p *PersonApi) GetById(c *gin.Context) {
	id := c.Param("id")
	people, errx := p.service.GetById(id)
	if errx != nil {
		c.JSON(400, gin.H{"error": errx.Cause().Error()})
		return
	}
	if people.Id == "" {
		c.JSON(404, gin.H{})
		return
	}
	c.JSON(http.StatusOK, people)
}

func (p *PersonApi) Post(c *gin.Context) {
	var createPerson model.CreatePerson
	if err := c.Bind(&createPerson); err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	person, err := p.service.Post(createPerson)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Cause().Error()})
		return
	}

	c.JSON(http.StatusOK, person)
}
