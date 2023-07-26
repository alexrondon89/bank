package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type dbConnection struct {
	conn *sqlx.DB
}

func CreateConnection() *dbConnection {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	if err != nil {
		log.Fatalln("error creating db connection: ", err)
	}

	return &dbConnection{
		conn: db,
	}
}
