package api

import (
	"github.com/Tonioou/go-api-test/dao"
	"github.com/gin-gonic/gin"
)

type PersonApi struct {
	dao dao.Database
}

func NewPersonApi() *PersonApi {
	return &PersonApi{
		dao: dao.NewInMemDatabase(),
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
