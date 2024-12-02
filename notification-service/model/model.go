package model

type NotificationRequest struct {
	Recipient string `json:"recipient"`
	OrderId   string `json:"order_id"`
	Message   string `json:"message"`
}
