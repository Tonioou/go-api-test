package api

import (
	"io"
	"net/http"
)

type HelloApi struct{}

func NewHelloApi() *HelloApi {
	return &HelloApi{}
}

func (hp *HelloApi) Register() {
	http.HandleFunc("/", hp.Hello)
}

func (hp *HelloApi) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"hello": "world"}`)
}
