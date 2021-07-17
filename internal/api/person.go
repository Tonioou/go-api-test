package api

import (
	"net/http"

	"github.com/Tonioou/go-api-test/internal/model"
	"github.com/Tonioou/go-api-test/internal/service"
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
	e.GET("/person", p.Get)
	e.POST("/person", p.Post)
}

func (p *PersonApi) Get(c *gin.Context) {
	people, err := p.service.Get()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Cause().Error()})
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
