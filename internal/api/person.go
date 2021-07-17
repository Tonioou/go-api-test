package api

import (
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
	// hp.dao.Add()
	// result := hp.dao.Get()
	// people := result.([]*model.Person)
	// json.NewEncoder(w).Encode(people)
}

func (p *PersonApi) Post(c *gin.Context) {

}
