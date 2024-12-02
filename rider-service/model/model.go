package model

type BaseData struct {
	Id   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type Rider struct {
	BaseData
}

type GetRidersResponse struct {
	Riders []Rider `json:"rider"`
}

func (resp *GetRidersResponse) AddRiders(riders []*Rider) {
	for _, r := range riders {
		resp.Riders = append(resp.Riders, *r)
	}
}

type Order struct {
	Id           uint   `gorm:"primaryKey" json:"id"`
	RestaurantId string `json:"restaurant_id"`
	RiderId      string `json:"rider_id"`
	Items        string `gorm:"default:[]"`
	Status       string `json:"status"`
}

type RiderUpdateOrderRequest struct {
	OrderId string `json:"order_id"`
	RiderId string `json:"rider_id"`
}

type NotificationRequest struct {
	Recipient string `json:"recipient"`
	OrderId   string `json:"order_id"`
	Message   string `json:"message"`
}

type NotificationResponse struct {
	Status string `json:"status"`
}
