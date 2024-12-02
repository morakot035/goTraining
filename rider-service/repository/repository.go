package repository

import (
	"goTraining/rider-service/config"
	"goTraining/rider-service/model"
)

func FindAllRider() ([]*model.Rider, error) {
	var db = config.Database()
	var riders []*model.Rider
	dbResponse := db.Find(&riders)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return riders, nil
}

func UpdateOrderStatus(request model.RiderUpdateOrderRequest, status string) (*model.Order, error) {
	var db = config.Database()
	var order *model.Order
	dbResponse := db.Where("id = ?", request.OrderId).Find(&order)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	order.Status = status
	order.RiderId = request.RiderId
	dbResponse = db.Save(&order)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return order, nil
}
