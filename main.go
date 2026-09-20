package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"lpnapi/repository"
	"lpnapi/router"
	"lpnapi/service"
	"net/http"
	"os"
)

func main() {
	log.Println("Starting server...")
	dbUrl := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	gameRepository := &repository.GameRepository{DB: pool}
	gameService := &service.GameService{Repo: gameRepository}
	r := router.NewRouter(gameService)
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
