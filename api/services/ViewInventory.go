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
type ViewInventoryRequest struct {
	Category string `json:"category"`
	SellerID string `json:"seller_id"`
}

// Struct for View Inventory response
type ViewInventoryResponse struct {
	Inventory []models.Inventory `json:"inventory"`
}

//ViewInventory is service to register particular face using image
func ViewInventory(db *sql.DB, config models.Config, w http.ResponseWriter, r *http.Request) {

	var req ViewInventoryRequest
	var resp ViewInventoryResponse

	w.Header().Set("Content-Type", "text/html")

	// request sanity check
	req, err := ViewInventorySanityCheck(r)
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

	inventory, err := repository.FetchInventory(req.SellerID, req.Category, tx)

	if err != nil {
		tx.Rollback()

		log.Println("DB Error - Fetch Inventory : ", err.Error())

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

	resp.Inventory = inventory

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&resp)

}

func ViewInventorySanityCheck(r *http.Request) (req ViewInventoryRequest, err error) {
	// parse seller_id parameter
	sellerIDs, ok := r.URL.Query()["seller_id"]

	// return error if param seller_id missing
	if !ok || len(sellerIDs[0]) < 1 {
		log.Println("Url Param 'seller_id' is missing")

		err = errors.New("seller_id is null - please refer to request format in documentation")
		return
	}

	// parse category parameter
	categories, ok := r.URL.Query()["category"]

	// return error if param category missing
	if !ok || len(categories[0]) < 1 {
		log.Println("Url Param 'category' is missing")

		err = errors.New("category is null - please refer to request format in documentation")
		return
	}

	req = ViewInventoryRequest{
		SellerID: sellerIDs[0],
		Category: categories[0]}

	log.Println("request : ", req)
	return
}
