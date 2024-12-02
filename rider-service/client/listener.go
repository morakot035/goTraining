package client

import (
	"context"
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
)

var topic = "rider"
var brokers = []string{"localhost:9093"}
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
		consumer.Close() // Close Consumer
	}()
	fmt.Println("Start Listening to Notification")
	for {
		select {
		case <-done:
			log.Info("Done Listening for Notification")
			return
		default:
			// Listening to Message
			msg, err := consumer.ReadMessage(context.Background())
			if err != nil {
				log.Info(fmt.Sprintf("Error reading notification %v", err))
				continue
			}
			if len(string(msg.Value)) > 0 {
				log.Info(fmt.Sprintf("Received Notification: %s\n", string(msg.Value)))
				consumer.CommitMessages(context.Background(), msg)
			}
		}
	}
}
