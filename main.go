package main

import (
	"lpnapi/router"
	"net/http"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8080", r)
}
