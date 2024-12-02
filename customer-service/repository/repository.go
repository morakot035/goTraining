package repository

import (
	"encoding/json"

	"goTraining/customer-service/config"
	"goTraining/customer-service/model"

	"gorm.io/gorm"
)

func CreateNewOrder(request model.PlaceOrderResquest) (*model.Order, error) {
	var db = config.Database()

	itemsData, err := json.Marshal(request.Items)
	if err != nil {
		return &model.Order{}, err
	}

	order := model.Order{
		RestaurantId: request.RestaurantId,
		Items:        string(itemsData),
		Status:       "created",
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&order)
		return nil
	})
	if err != nil {
		return &model.Order{}, err
	}
	return &order, nil
}
