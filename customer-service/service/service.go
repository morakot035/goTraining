package service

import (
	"fmt"
	"goTraining/customer-service/model"
	"goTraining/customer-service/repository"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func PlaceOrder(ctx echo.Context) error {
	var request model.PlaceOrderResquest
	err := ctx.Bind(&request)
	if err != nil {
		log.Error(err)
		return echo.ErrBadRequest
	}

	order, err := repository.CreateNewOrder(request)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	return ctx.JSON(http.StatusOK, model.OrderStatusResponse{OrderId: fmt.Sprintf("%v", order.Id), Status: order.Status})
}
