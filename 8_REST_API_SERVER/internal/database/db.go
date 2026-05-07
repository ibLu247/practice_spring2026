package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	var err error

	DB, err = pgxpool.New(
		context.Background(),
		"postgres://postgres:pass@localhost:5432/postgres",
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")
}
