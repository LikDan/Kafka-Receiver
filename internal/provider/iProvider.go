package provider

import (
	"context"
	"receiver/internal/entities"
)

type IProvider interface {
	WriteAnswer(ctx context.Context, message entities.AnsweredMessage) error
	ReadMessage(ctx context.Context) (entities.Message, error)
}
