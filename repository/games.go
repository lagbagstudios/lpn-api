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

func (r *GameRepository) FetchLPN(gameCode int16) (int16, error) {
	var lpn int16
	err := r.DB.QueryRow(context.Background(), `select lpn from games where code=$1`, gameCode).Scan(&lpn)
	if err != nil {
		log.Printf("Error reading from DB: %s", err.Error())
		return 0, err
	}

	return lpn, nil
}

func (r *GameRepository) CreateGame(gameCode int16) (model.Game, error) {
	var lpn int16
	err := r.DB.QueryRow(context.Background(), `insert into games(code, lpn) values ($1, $2) on conflict do nothing returning lpn`, gameCode, 0).Scan(&lpn)
	if err != nil {
		return model.Game{}, err
	}
	return model.Game{Code: gameCode, LPN: lpn}, nil
}

func (r *GameRepository) UpdateLPN(game *model.Game) (int16, error) {
	var updatedLpn int16
	err := r.DB.QueryRow(context.Background(), `update games set lpn=$2 where code=$1 returning lpn`, game.Code, game.LPN).Scan(&updatedLpn)
	if err != nil {
		log.Printf("Error updating LPN: %s", err.Error())
		return 0, err
	}
	return updatedLpn, nil
}
