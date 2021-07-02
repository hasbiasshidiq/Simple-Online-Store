package models

// OrderItems struct
type OrderItems struct {
	OrderItemID int    `json:"order_item_id"`
	OrderID     int    `json:"order_id"`
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Quantity    int    `json:"quantity"`
	TotalPrice  int    `json:"total_price"`
}
