package repository

import (
	"database/sql"
	models "store/api/models"
)

type OrderItems struct {
	*models.OrderItems
}

func (orderItems *OrderItems) InsertOrderItems(tx *sql.Tx) (lastInsertID int, err error) {

	// INSERT INTO orders table
	sqlStatement := `INSERT INTO order_items
		(order_id, product_id, quantity, total_price) 
	VALUES 
		($1, $2, $3, $4)
	RETURNING order_id;`

	row := tx.QueryRow(sqlStatement, orderItems.OrderID, orderItems.ProductID, orderItems.Quantity, orderItems.TotalPrice)
	err = row.Scan(&lastInsertID)

	return
}
