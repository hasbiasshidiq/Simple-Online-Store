package repository

import (
	"database/sql"
	models "store/api/models"
	"time"
)

type Order struct {
	*models.Order
}

func (order *Order) InsertOrder(tx *sql.Tx) (lastInsertID int, err error) {

	// INSERT INTO orders table
	sqlStatement := `INSERT INTO orders
		(seller_id, customer_id, order_status, order_time) 
	VALUES 
		($1, $2, $3, $4)
	RETURNING order_id;`

	row := tx.QueryRow(sqlStatement, order.SellerID, order.CustomerID, order.OrderStatus, time.Now())
	err = row.Scan(&lastInsertID)

	return
}
