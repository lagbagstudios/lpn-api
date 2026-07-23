package router

import (
	"lpnapi/service"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	game := service.GameHandler{}
	r := mux.NewRouter()
	r.HandleFunc("/game/{code}", game.GetLPN)
	return r
}
