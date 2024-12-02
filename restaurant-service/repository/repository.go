package repository

import (
	"encoding/json"

	"goTraining/restaurant-service/config"
	"goTraining/restaurant-service/model"

	"github.com/labstack/gommon/log"
)

func FindAllRestaurant() (*[]model.Restaurant, error) {
	var db = config.Database()
	var restaurantsDB *[]model.RestaurantDB
	dbResponse := db.Find(&restaurantsDB)

	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var restaurants []model.Restaurant
	for _, r := range *restaurantsDB {
		var menu []model.Menu
		err := json.Unmarshal([]byte(r.Menu), &menu)
		if err != nil {
			return nil, err
		}
		restaurants = append(restaurants, model.Restaurant{BaseData: r.BaseData, Menu: menu})
	}
	return &restaurants, nil
}

func FindRestaurantById(id uint) (*model.Restaurant, error) {
	var db = config.Database()
	var restaurant *model.RestaurantDB
	dbResponse := db.Find(&restaurant, "id = ?", id)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	var menu []model.Menu
	err := json.Unmarshal([]byte(restaurant.Menu), &menu)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &model.Restaurant{BaseData: restaurant.BaseData, Menu: menu}, nil
}

func AcceptingOrder(request model.AcceptOrderRequest) (*model.Order, error) {
	var db = config.Database()
	var order *model.Order
	dbResponse := db.Where("id = ? AND restaurantId = ?", request.OrderId, request.RestaurantId).Find(&order)
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	return UpdateOrderStatus(order, "accepted")
}

func UpdateOrderStatus(order *model.Order, status string) (*model.Order, error) {
	var db = config.Database()
	dbResponse := db.Model(&order).Update("status", "accepted")
	if dbResponse.Error != nil {
		return nil, dbResponse.Error
	}
	order.Status = "accepted"
	return order, nil
}
