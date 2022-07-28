package repository

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"receiver/internal/entities"
)

type Repository struct {
	answerWriter *kafka.Writer
}

func NewRepository(answerWriter *kafka.Writer) *Repository {
	return &Repository{answerWriter: answerWriter}
}

func (r *Repository) WriteAnswer(ctx context.Context, message entities.AnsweredMessage) error {
	bytesMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := kafka.Message{Value: bytesMessage}
	return r.answerWriter.WriteMessages(ctx, kafkaMessage)
}
