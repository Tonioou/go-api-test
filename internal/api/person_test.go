package api_test

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"

// 	"github.com/Tonioou/go-person-crud/api"
// )

// func TestHelloApi(t *testing.T) {
// 	personApi := &api.PersonApi{}
// 	req := httptest.NewRequest("GET", "/", nil)
// 	w := httptest.NewRecorder()
// 	personApi.Person(w, req)
// 	fmt.Println(w.Body.String())
// 	//assert.Equal(t, w.Body.String(), `[{"Email":"dorothy@aol.com","Name":"Dorothy","Age":53},{"Email":"joe@aol.com","Name":"Joe","Age":30},{"Email":"lucy@aol.com","Name":"Lucy","Age":35},{"Email":"tariq@aol.com","Name":"Tariq","Age":21}]`)
// 	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
// 	assert.NotEqual(t, http.StatusBadGateway, w.Result().StatusCode)
// }
