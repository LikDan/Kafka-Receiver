package controller

import (
	"context"
	"receiver/internal/entities"
	"receiver/internal/provider"
)

type Controller struct {
	provider provider.IProvider
}

func NewController(provider provider.IProvider) *Controller {
	return &Controller{provider: provider}
}

func (c *Controller) Answer(ctx context.Context, message entities.Message) error {
	answeredMessage := entities.AnsweredMessage{
		Id:      message.Id,
		Message: message.Message,
		Answer:  "Answer",
	}

	if err := c.provider.WriteAnswer(ctx, answeredMessage); err != nil {
		return err
	}

	return nil
}

func (c *Controller) Read(ctx context.Context) (entities.Message, error) {
	return c.provider.ReadMessage(ctx)
}
