package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool

func InitDb() error {
	log.Println("Initialising database")

	config, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("Failed to parse database config: ", err)
		return err
	}

	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		// do something with every new connection
		log.Println("new db connection")
		return nil
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Println("Failed to connect to database: ", err)
		return err
	}
	db = pool

	return nil
}

func Close() {
	if db != nil {
		db.Close()
	}
}

func Pool() *pgxpool.Pool {
	return db
}
