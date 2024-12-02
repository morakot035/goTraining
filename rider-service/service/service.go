package service

import (
	"fmt"
	"net/http"

	"goTraining/rider-service/client"
	"goTraining/rider-service/model"
	"goTraining/rider-service/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetRiders(ctx echo.Context) error {
	log.Info("Get All Riders Data")

	riders, err := repository.FindAllRider()

	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	var response model.GetRidersResponse
	response.AddRiders(riders)
	return ctx.JSON(http.StatusOK, response)
}

func PickUpOrder(ctx echo.Context) error {
	var request model.RiderUpdateOrderRequest
	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}
	order, err := repository.UpdateOrderStatus(request, "picked_up")
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	notificationRequest := model.NotificationRequest{
		Recipient: "customer",
		OrderId:   request.OrderId,
		Message:   fmt.Sprintf("Order %s had been picked up by rider %s.", request.OrderId, request.RiderId),
	}
	err = client.SendNotification(notificationRequest)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"status": order.Status})
}

func DeliverOrder(ctx echo.Context) error {
	var request model.RiderUpdateOrderRequest
	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}
	order, err := repository.UpdateOrderStatus(request, "delivered")
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	notificationRequest := model.NotificationRequest{
		Recipient: "customer",
		OrderId:   request.OrderId,
		Message:   fmt.Sprintf("Order %s had been delivered by rider %s.", request.OrderId, request.RiderId),
	}
	err = client.SendNotification(notificationRequest)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	return ctx.JSON(http.StatusOK, map[string]string{"status": order.Status})
}
