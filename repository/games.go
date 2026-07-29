package repository

import (
	"context"
	"log"
	"lpnapi/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GameRepository struct {
	DB *pgxpool.Pool
}

func (r *GameRepository) FetchLPN(gameCode int16) int16 {
	var lpn int16
	err := r.DB.QueryRow(context.Background(), `select lpn from games where code=$1`, gameCode).Scan(&lpn)
	if err != nil {
		log.Printf("Error reading from DB: %s", err.Error())
		return -1
	}

	return lpn
}

func (r *GameRepository) CreateGame(gameCode int16) model.Game {
	var lpn int16
	err := r.DB.QueryRow(context.Background(), `insert into games(code, lpn) values ($1, $2) on conflict do nothing returning lpn`, gameCode, 0).Scan(&lpn)
	if err != nil {
		return model.Game{Code: -1, LPN: -1}
	}
	return model.Game{Code: gameCode, LPN: lpn}
}

func (r *GameRepository) UpdateLPN(gameCode, lpn int16) int16 {
	var updatedLpn int16
	err := r.DB.QueryRow(context.Background(), `update games set lpn=$2 where code=$1 returning lpn`, gameCode, lpn).Scan(&updatedLpn)
	if err != nil {
		log.Printf("Error updating LPN: %s", err.Error())
		return -1
	}
	return updatedLpn
}
