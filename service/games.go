package service

import (
	"encoding/json"
	"fmt"
	"lpnapi/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type gameStore interface {
	Get(code int) model.Game
}

type GameHandler struct{}

func (g *GameHandler) GetLPN(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	codeStr, ok := vars["code"]
	if !ok {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	// TODO: Fetch from DB using code

	code, err := strconv.Atoi(codeStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Code must be numeric: %s", err.Error()), http.StatusBadRequest)
		return
	}

	responseData := model.Game{Code: code, LPN: 123}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte("123"))
}
