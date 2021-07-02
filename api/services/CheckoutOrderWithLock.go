package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	models "store/api/models"
	repository "store/api/repository"

	"net/http"
)

//CheckoutOrderWithLock is service to register particular face using image
func CheckoutOrderWithLock(db *sql.DB, config models.Config, w http.ResponseWriter, r *http.Request) {

	var req CheckoutOrderRequest
	var resp CheckoutOrderResponse

	w.Header().Set("Content-Type", "text/html")

	// request sanity check
	req, err := CheckoutOrderSanityCheck(r)
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

	order := &models.Order{
		CustomerID:  *req.CustomerID,
		SellerID:    *req.SellerID,
		OrderStatus: *req.OrderStatus,
	}

	orderRepo := &repository.Order{order}
	ID, err := orderRepo.InsertOrder(tx)
	if err != nil {
		log.Println("DB Error - Insert Order : ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	order.OrderID = ID
	orderItemsList := req.OrderItems

	failedOrderItems, httpStatusError, httpErrorMessage := HandleOrderItemsWithLock(order, orderItemsList, tx)
	if httpStatusError != nil {
		w.WriteHeader(*httpStatusError)
		w.Write([]byte(httpErrorMessage))
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println("DB Error - Problem encountered during db commit : ", err.Error())

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	// return list of failed order items (order items where inventory quantity less than ordered quantity)
	if len(failedOrderItems) > 0 {

		resp.FailedOrderItems = failedOrderItems

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(&resp)
		return
	}

	log.Println("Success")

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Success"))

}

// Handle Order Items Flow :
// 1. Get item available quantity from inventory table
// 2. Check if ordered items quantity > item available quantity
// 3. Update Inventory if step(2) returns true
// 4. Insert Order Items
func HandleOrderItemsWithLock(order *models.Order, orderItemsList []models.OrderItems, tx *sql.Tx) (
	failedOrderItems []models.OrderItems,
	httpErrorStatus *int,
	httpErrorMessage string) {

	for i := range orderItemsList {

		orderItemsList[i].OrderID = order.OrderID

		// check product quantity on inventory table
		quantity, err := repository.FetchQuantityWithLock(order.SellerID, orderItemsList[i].ProductID, tx)
		if err == sql.ErrNoRows {
			tx.Rollback()

			log.Println("Some products is not registered or sellerid-productid mismatch")

			httpStatus := http.StatusNotAcceptable
			httpErrorStatus = &httpStatus
			httpErrorMessage = "Some products is not registered or sellerid-productid mismatch"
			break
		}

		// handle general db error
		if err != nil {
			log.Println("DB Error - Check Quantity : ", err.Error())

			httpStatus := http.StatusInternalServerError
			httpErrorStatus = &httpStatus
			httpErrorMessage = "Internal Server Error"
			break
		}

		// if quantity in inventory is less than quantity of ordered items, then skip to the next ordered items
		if quantity < orderItemsList[i].Quantity {
			failedOrderItems = append(failedOrderItems, orderItemsList[i])
			continue
		}

		// update inventory
		err = repository.DecreaseInventory(order.SellerID, orderItemsList[i].ProductID, orderItemsList[i].Quantity, tx)
		if err != nil {
			tx.Rollback()

			log.Println("DB Error - Update Inventory : ", err.Error())

			httpStatus := http.StatusInternalServerError
			httpErrorStatus = &httpStatus
			httpErrorMessage = "Internal Server Error"
			break
		}

		// insert order items for logging purpose
		orderItemsRepo := &repository.OrderItems{&orderItemsList[i]}
		_, err = orderItemsRepo.InsertOrderItems(tx)
		if err != nil {
			tx.Rollback()

			log.Println("DB Error - Insert Order Item : ", err.Error())

			httpStatus := http.StatusInternalServerError
			httpErrorStatus = &httpStatus
			httpErrorMessage = "Internal Server Error"
			break
		}
	}
	return
}
