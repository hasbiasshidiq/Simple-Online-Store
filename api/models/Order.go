package models

// Order struct is model inventory table
type Order struct {
	OrderID     int    `json:"order_id"`
	CustomerID  string `json:"customer_id"`
	SellerID    string `json:"seller_id"`
	OrderStatus string `json:"order_status"`
}
