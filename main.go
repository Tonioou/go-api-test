package main

import (
	"net/http"

	"github.com/Tonioou/go-api-test/api"
)

func main() {
	api.NewHelloApi().Register()
	http.ListenAndServe(":8080", nil)
}
