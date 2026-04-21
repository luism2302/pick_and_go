package main

import (
	"context"
	"log"
	"os"
	"pick_and_go/mlb"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Couldn't load .env: %s", err)
	}

	dbString := os.Getenv("DB_STRING")
	if dbString == "" {
		log.Fatal("Couldn't find DB_STRING in .env")
	}

	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		log.Fatalf("Couldn't connect to db: %s", err)
	}
	defer conn.Close(context.Background())

	client := mlb.NewSportClient(conn)

	if err := client.ResetResults(); err != nil {
		log.Fatal(err)
	}
	if err := client.UpdateResults(); err != nil {
		log.Fatal(err)
	}

}
