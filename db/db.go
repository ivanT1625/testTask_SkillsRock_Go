package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
)

var DB *pgx.Conn

func InitDB(connString string) {
	var err error
	DB, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Не удается подключить к базе данных: %v\n", err)
	}

	if DB == nil {
		log.Fatalf("DB = nil")
	}

	log.Println("Подключен к базе данных")
}
