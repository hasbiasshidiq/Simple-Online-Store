package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	models "store/api/models"
	repository "store/api/repository"

	"net/http"
)

// Struct for View Inventory request
type UpdateInventoryRequest struct {
	SellerID    *string `json:"seller_id"`
	ProductID   *int    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Quantity    *int    `json:"quantity"`
}

// // Struct for View Inventory response
// type UpdateInventoryResponse struct {
// 	Inventory []models.Inventory `json:"inventory"`
// }

//UpdateInventory is service to register particular face using image
func UpdateInventory(db *sql.DB, config models.Config, w http.ResponseWriter, r *http.Request) {

	var req UpdateInventoryRequest
	// var resp UpdateInventoryResponse

	w.Header().Set("Content-Type", "text/html")

	// request sanity check
	req, err := UpdateInventorySanityCheck(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("DB Error - Problem encountered during tx creation : ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = repository.UpdateInventory(*req.SellerID, *req.ProductID, *req.Quantity, tx)

	if err == sql.ErrNoRows {
		log.Println("Inventory Data is not found or sellerid-productid mismatch")

		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Inventory Data is not found or sellerid-productid mismatch"))
		return
	}
	if err != nil {
		tx.Rollback()

		log.Println("DB Error - Update Inventory : ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("DB Error - Problem encountered during db commit : ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Success"))
}

func UpdateInventorySanityCheck(r *http.Request) (req UpdateInventoryRequest, err error) {

	errDecode := json.NewDecoder(r.Body).Decode(&req)
	if errDecode != nil {
		log.Println("error parsing request : ", errDecode)
		err = errors.New("request format malformed - please refer to request format in documentation")
		return
	}

	if req.SellerID == nil {
		log.Println("SellerID is null")
		err = errors.New("seller_id is null - please refer to request format in documentation")
		return
	}

	if req.ProductID == nil {
		log.Println("ProductID is null")
		err = errors.New("product_id is null - please refer to request format in documentation")
		return
	}

	if req.Quantity == nil {
		log.Println("Quantity is null")
		err = errors.New("quantity is null - please refer to request format in documentation")
		return
	}
	return
}
