package service

import (
	"context"
	"fmt"
	"net/http"

	"goTraining/notification-service/config"
	"goTraining/notification-service/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
)

func NotificationSend(ctx echo.Context) error {
	var request model.NotificationRequest
	err := ctx.Bind(&request)
	if err != nil || len(request.Recipient) == 0 {
		log.Error(err)
		return echo.ErrBadRequest
	}
	writer := config.GetWriter(request.Recipient)
	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(fmt.Sprintf("Order %s Message: %s", request.OrderId, request.Message)),
		},
	)
	if err != nil {
		log.Error(err)
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return ctx.JSON(http.StatusOK, data)
	}
	writer.Close()
	data := map[string]string{
		"status": "send",
	}
	return ctx.JSON(http.StatusOK, data)
}
