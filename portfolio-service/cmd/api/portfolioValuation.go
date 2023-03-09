package main

import (
	"portfolio-service/data"
	"time"
)

func GetQuantityByDate(date time.Time, symbol string, transactionHistory []data.Transaction) int {
	var stockQuantity int
	for _, transaction := range transactionHistory {
		if date.After(transaction.Date) && transaction.Symbol == symbol {
			if transaction.Type == "BUY" {
				stockQuantity = stockQuantity + transaction.Quantity
			} else if transaction.Type == "SELL" {
				stockQuantity = stockQuantity - transaction.Quantity
			}
		}
	}

	return stockQuantity
}
