package data

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

func ConnectToPostgresSQL() (PostgreSQL, error) {
	fmt.Println("Connectiong to postgreSQL...")

	dsn := os.Getenv("DSN")
	fmt.Printf("Connectiong to postgreSQL with dsn %s\n", dsn)
	var counts int64
	var postgreSQL PostgreSQL

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Postgres not ready yet...")
			counts++
		} else {
			log.Println("Connected to postgres!")
			postgreSQL.db = connection
			return postgreSQL, nil
		}

		if counts > 10 {
			log.Println(err)
			return postgreSQL, errors.New("unable to connect to PostgreSQL timout")
		}

		log.Println("Backing of for two seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return db, nil
	}

	return db, nil
}

type PostgreSQL struct {
	db *sql.DB
}
