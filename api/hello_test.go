package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Tonioou/go-api-test/api"
)

func TestHelloApi(t *testing.T) {
	helloApi := &api.HelloApi{}
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	helloApi.Hello(w, req)
	assert.Equal(t, w.Body.String(), `{"hello": "world"}`)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.NotEqual(t, http.StatusBadGateway, w.Result().StatusCode)
}
