package router

import (
	"lpnapi/service"

	"github.com/gorilla/mux"
)

func NewRouter(gameService *service.GameService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/game/{code}", gameService.GetLPN)
	return r
}
