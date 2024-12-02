package client

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
)

var topic = "restaurant"
var brokers = []string{"localhost:8080"}
var groupId = "food-delivery"

func createConsumer() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupId,
		Topic:   topic,
	})
}

func ListenForNotification(done <-chan bool) {
	consumer := createConsumer()
	defer func() {
		log.Info("Close Consumer")
		consumer.Close()
	}()
	fmt.Println("Start Listening to Notification")

	for {
		select {
		case <-done:
			log.Info("Done Listening for Noti")
			return
		default:
			msg, err := consumer.ReadMessage(context.Background())
			if err != nil {
				log.Info(fmt.Sprintf("Error reading %v", err))
				continue
			}
			if len(string(msg.Value)) > 0 {
				log.Info(fmt.Sprintf("Received Noti: %s\n", string(msg.Value)))
				consumer.CommitMessages(context.Background(), msg)
			}
		}
	}
}
