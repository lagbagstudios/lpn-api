package router

import (
	"lpnapi/service"
	"net/http"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewRouter(gameService *service.GameService) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/game/{code}", gameService.GetLPN).Methods("GET", "OPTIONS")
	r.HandleFunc("/game/{code}", gameService.UpdateLPN).Methods("PUT", "OPTIONS")
	r.HandleFunc("/game", gameService.CreateGame).Methods("POST", "OPTIONS")
	r.Use(corsMiddleware)
	return r
}
