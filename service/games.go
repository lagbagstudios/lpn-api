package service

import (
	"encoding/json"
	"fmt"
	"lpnapi/model"
	"lpnapi/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type gameStore interface {
	Get(code int) model.Game
}

type GameService struct {
	Repo *repository.GameRepository
}

func (g *GameService) GetLPN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	codeStr, ok := vars["code"]
	if !ok {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	code64, err := strconv.Atoi(codeStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code must be numeric: %s", err.Error()), http.StatusBadRequest)
		return
	}
	code := int16(code64)

	lpn := g.Repo.FetchLPN(code)
	if lpn < 0 {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	responseData := model.Game{Code: code, LPN: lpn}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (g *GameService) CreateGame(w http.ResponseWriter, r *http.Request) {
	var request model.CreateGameRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `Invalid payload, expected type is {"code": int16}`, http.StatusBadRequest)
	}

	gameResponse := g.Repo.CreateGame(request.Code)
	if gameResponse.Code < 0 {
		http.Error(w, "Unexpected error creating new game", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(gameResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (g *GameService) UpdateLPN(w http.ResponseWriter, r *http.Request) {
	var request model.UpdateLPNRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, `Invalid payload, expected {"code": int16, "lpn": int16}`, http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	codeStr, ok := vars["code"]
	if !ok {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	code64, err := strconv.Atoi(codeStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code must be numeric: %s", err.Error()), http.StatusBadRequest)
		return
	}
	code := int16(code64)

	gameResponse := g.Repo.UpdateLPN(&model.Game{Code: code, LPN: request.LPN})
	if gameResponse < 0 {
		http.Error(w, "Unexpected error updating LPN", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(gameResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
