package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgreSQL struct {
	db *sql.DB
}

const dbTimeout = time.Second * 3

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

func (p *PostgreSQL) getPortfolioByUserId(userId int) (*Portfolio, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT symbol, sum(price*quantity)/sum(quantity), SUM (quantity), sum(price*quantity)
	FROM public.transactions 
	where user_id = $1
	GROUP BY symbol`

	rows, err := p.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	var portfolio *Portfolio

	for rows.Next() {
		var position Position
		err := rows.Scan(
			&position.Symbol,
			&position.Quantity,
			&position.Price,
			&position.Value,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		portfolio.Positions = append(portfolio.Positions, position)
	}

	return portfolio, nil
}
