package main

import (
	"context"
	"log"
	"lpnapi/repository"
	"lpnapi/router"
	"lpnapi/service"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	connectionString := "postgresql://lpnuser:lpnpassword@db:5432/lpn"
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	gameRepository := &repository.GameRepository{DB: pool}
	gameService := &service.GameService{Repo: gameRepository}
	r := router.NewRouter(gameService)
	http.ListenAndServe(":8080", r)
}
