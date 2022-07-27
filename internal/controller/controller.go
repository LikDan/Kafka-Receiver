package controller

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"receiver/internal/entities"
)

type Controller struct {
	connection *kafka.Conn
}

func NewController(connection *kafka.Conn) *Controller {
	return &Controller{connection: connection}
}

func (c *Controller) Answer(_ context.Context, message entities.Message) error {
	answeredMessage := entities.AnsweredMessage{
		Id:      message.Id,
		Message: message.Message,
		Answer:  "Answer",
	}

	bytesMessage, err := json.Marshal(answeredMessage)
	if err != nil {
		return err
	}

	kafkaMessage := kafka.Message{Value: bytesMessage}
	if _, err = c.connection.WriteMessages(kafkaMessage); err != nil {
		return err
	}

	return nil
}
