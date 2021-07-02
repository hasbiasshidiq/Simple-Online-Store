package repository

import (
	"database/sql"
	"fmt"
	"log"
	models "store/api/models"
)

// type Inventory struct {
// 	*models.Inventory
// }

func FetchInventory(SellerID, Category string, tx *sql.Tx) (inventories []models.Inventory, err error) {

	var sqlStatementTemplate, sqlStatement string

	// %s is placeholder for where clause
	sqlStatementTemplate = `
		SELECT 
			inventory.seller_id, 
			sellers.seller_name,
			inventory.product_id,
			products.product_name,
			products.category,
			products.price,
			inventory.quantity
		FROM 
			inventory
		INNER JOIN sellers
			ON inventory.seller_id = sellers.seller_id
		INNER JOIN products
			ON inventory.product_id = products.product_id
		%s
		ORDER BY (inventory.seller_id, products.category);`

	if (SellerID != "all") && (Category != "all") {
		sqlStatement = fmt.Sprintf(sqlStatementTemplate, "WHERE inventory.seller_id = '%s' AND products.category = '%s'")
		sqlStatement = fmt.Sprintf(sqlStatement, SellerID, Category)

		goto query
	}

	if SellerID != "all" {
		sqlStatement = fmt.Sprintf(sqlStatementTemplate, "WHERE inventory.seller_id = '%s'")
		sqlStatement = fmt.Sprintf(sqlStatement, SellerID)

		goto query
	}

	if Category != "all" {
		sqlStatement = fmt.Sprintf(sqlStatementTemplate, "WHERE products.category = '%s'")
		sqlStatement = fmt.Sprintf(sqlStatement, Category)

		goto query
	}

	// placeholder is filled with nothing so there are no where clause
	sqlStatement = fmt.Sprintf(sqlStatementTemplate, "")

query:
	rows, err := tx.Query(sqlStatement)

	if err != nil {
		// handle this error better than this
		log.Println("error : ", err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {

		var inventory models.Inventory

		err = rows.Scan(
			&inventory.SellerID,
			&inventory.SellerName,
			&inventory.ProductID,
			&inventory.ProductName,
			&inventory.Category,
			&inventory.Price,
			&inventory.Quantity)

		if err != nil {
			log.Println("error in loop", err.Error())
			break
		}

		inventories = append(inventories, inventory)

	}
	// get any error encountered during iteration
	err = rows.Err()

	return
}

// get product quantity on inventory table by sellerID and productID
func FetchQuantity(SellerID string, ProductID int, tx *sql.Tx) (Quantity int, err error) {
	sqlStatement := `
	SELECT quantity FROM inventory
	WHERE seller_id = $1 AND product_id = $2;`

	row := tx.QueryRow(sqlStatement, SellerID, ProductID)
	err = row.Scan(&Quantity)
	return
}

// Fetch Quantity With Row Level Locking, Locked until transaction successfully commited or rollback
func FetchQuantityWithLock(SellerID string, ProductID int, tx *sql.Tx) (Quantity int, err error) {
	sqlStatement := `
	SELECT quantity FROM inventory
	WHERE seller_id = $1 AND product_id = $2 
	FOR UPDATE;`

	row := tx.QueryRow(sqlStatement, SellerID, ProductID)
	err = row.Scan(&Quantity)
	return
}

// UpdateInventory update quantity of existing product
func UpdateInventory(SellerID string, ProductID int, Quantity int, tx *sql.Tx) (err error) {
	sqlStatement := "UPDATE inventory SET quantity = $1 WHERE seller_id = $2 AND product_id = $3"

	result, err := tx.Exec(sqlStatement, Quantity, SellerID, ProductID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if rowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return
}

// DecreaseInventory update quantity of existing product
func DecreaseInventory(SellerID string, ProductID int, QuantityChange int, tx *sql.Tx) (err error) {
	sqlStatement := "UPDATE inventory SET quantity = quantity - $1 WHERE seller_id = $2 AND product_id = $3"

	result, err := tx.Exec(sqlStatement, QuantityChange, SellerID, ProductID)
	if err != nil {
		log.Println(err.Error())
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return
	}

	if rowsAffected == 0 {
		err = sql.ErrNoRows
	}

	return
}
