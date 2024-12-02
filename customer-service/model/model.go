package model

type Order struct {
	Id           uint   `gorm:"primaryKey" json:"id"`
	RestaurantId string `json:"restaurant_id"`
	RiderId      string `json:"rider_id"`
	Items        string `gorm:"default:[]"`
	Status       string `json:"status"`
}

type PlaceOrderResquest struct {
	RestaurantId string `json:"restaurant_id"`
	Items        []Item `json:"items"`
}

type OrderStatusResponse struct {
	OrderId string `json:"order_id"`
	Status  string `json:"status"`
}

type Item struct {
	MenuId   string `json:"menu_id"`
	Quantity uint   `json:"quantity"`
}
