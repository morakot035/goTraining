package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goTraining/restaurant-service/model"
	"io"
	"net/http"

	"github.com/labstack/gommon/log"
)

const NotificationSendUrl = "http://localhost:8088/notification/send"

func SendNotification(request model.NotificationRequest) error {
	log.Info(fmt.Sprintf("Sending Request /notification/send: %v", request))
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return err
	}
	response, err := http.Post(NotificationSendUrl, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		message := fmt.Sprintf("Non-OK HTTP status: %d\n", response.StatusCode)
		return errors.New(message)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	var result model.NotificationResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Receive Response /notification/send: %v", result))
	return nil
}
