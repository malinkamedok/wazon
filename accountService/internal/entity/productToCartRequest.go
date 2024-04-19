package entity

type ProductToCartRequest struct {
	UserID    string `json:"userID"`
	ProductID string `json:"productID"`
}
