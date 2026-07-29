package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GameRepository struct {
	DB *pgxpool.Pool
}

func (r *GameRepository) FetchLPN(gameCode int16) int16 {
	var lpn int16
	err := r.DB.QueryRow(context.Background(), "select lpn from games where code=$1", gameCode).Scan(&lpn)
	if err != nil {
		log.Fatalf("Error reading from DB: %s", err.Error())
	}

	return int16(lpn)
}
