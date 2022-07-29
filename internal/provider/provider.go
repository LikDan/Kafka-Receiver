package provider

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"receiver/internal/entities"
)

type Provider struct {
	answerWriter  *kafka.Writer
	messageReader *kafka.Reader
}

func NewProvider(answerWriter *kafka.Writer, messageReader *kafka.Reader) *Provider {
	return &Provider{answerWriter: answerWriter, messageReader: messageReader}
}

func (r *Provider) WriteAnswer(ctx context.Context, message entities.AnsweredMessage) error {
	bytesMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	kafkaMessage := kafka.Message{Value: bytesMessage}
	return r.answerWriter.WriteMessages(ctx, kafkaMessage)
}

func (r *Provider) ReadMessage(ctx context.Context) (entities.Message, error) {
	kafkaMessage, err := r.messageReader.ReadMessage(ctx)
	if err != nil {
		return entities.Message{}, err
	}

	var message entities.Message
	if err = json.Unmarshal(kafkaMessage.Value, &message); err != nil {
		return entities.Message{}, err
	}

	return message, nil
}
