package models

// Inventory struct is model inventory table
type Inventory struct {
	SellerID    string `json:"seller_id"`
	SellerName  string `json:"seller_name"`
	ProductID   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Category    string `json:"category"`
	Quantity    string `json:"quantity"`
	Price       string `json:"price"`
}
