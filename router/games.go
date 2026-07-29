package router

import (
	"lpnapi/service"

	"github.com/gorilla/mux"
)

func NewRouter(gameService *service.GameService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/game/{code}", gameService.GetLPN).Methods("GET")
	r.HandleFunc("/game/{code}", gameService.UpdateLPN).Methods("PUT")
	r.HandleFunc("/game", gameService.CreateGame).Methods("POST")
	return r
}
