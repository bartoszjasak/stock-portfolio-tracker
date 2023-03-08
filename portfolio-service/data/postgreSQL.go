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

func (p *PostgreSQL) GetPortfolioByUserId(userId int) (Portfolio, error) {
	var portfolio Portfolio
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT stock_name, symbol, sum(price*quantity)/sum(quantity), SUM (quantity), sum(price*quantity)
	FROM public.transactions 
	where user_id = $1
	GROUP BY stock_name, symbol`

	rows, err := p.db.QueryContext(ctx, query, userId)
	if err != nil {
		return portfolio, err
	}

	for rows.Next() {
		var position Position
		err := rows.Scan(
			&position.StockName,
			&position.Symbol,
			&position.Price,
			&position.Quantity,
			&position.Value,
		)
		if err != nil {
			log.Fatal("Error scanning", err)
			return portfolio, err
		}

		portfolio.Positions = append(portfolio.Positions, position)
	}

	log.Printf("portfolio: %d", len(portfolio.Positions))
	return portfolio, nil
}

func (p *PostgreSQL) GetTransactionHostoryByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, type, stock_name, symbol, price, quantity, date, user_id
					  FROM public.transactions 
					  WHERE user_id = $1
						ORDER BY date ASC`

	rows, err := p.db.QueryContext(ctx, query, userId)
	if err != nil {
		log.Fatal("Query error", err)
		return transactions, err
	}

	for rows.Next() {
		var transaction Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.Type,
			&transaction.StockName,
			&transaction.Symbol,
			&transaction.Price,
			&transaction.Quantity,
			&transaction.Date,
			&transaction.UserId,
		)
		if err != nil {
			log.Fatal("Error scanning getTransactionHostoryByUserId response")
			return transactions, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}
