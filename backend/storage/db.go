package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectToDB() {
	connStr := "postgres://apixuser:apixpass@localhost:5432/apixchange"

	var err error

	DB, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	log.Println("Connected to database successfully!")
}

func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
