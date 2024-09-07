package models

type Session struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Total  string `json:"total"`
	// CreatedAt  time.DateTime `json:"created_at"`
	// ModifiedAt time.DateTime `json:"modified_at"`
}

type Item struct {
	ProductId int `json:"product_id"`
	SessionId int `json:"session_id"`
	Quantity  int `json:"quantity"`
}
