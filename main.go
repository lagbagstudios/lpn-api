package main

import (
	"context"
	"log"
	"lpnapi/repository"
	"lpnapi/router"
	"lpnapi/service"
	"net/http"
	"os"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("Starting server...")
	dbUrl := os.Getenv("DB_URL")

	migrateUrl := strings.Replace(dbUrl, "postgresql://", "postgres://", 1)
	m, err := migrate.New("file://migrations", migrateUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

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
