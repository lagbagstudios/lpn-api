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

	responseData := model.Game{Code: code, LPN: lpn}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
