package controller

import (
	"context"
	"receiver/internal/entities"
	"receiver/internal/repository"
)

type Controller struct {
	repository repository.IRepository
}

func NewController(repository repository.IRepository) *Controller {
	return &Controller{repository: repository}
}

func (c *Controller) Answer(ctx context.Context, message entities.Message) error {
	answeredMessage := entities.AnsweredMessage{
		Id:      message.Id,
		Message: message.Message,
		Answer:  "Answer",
	}

	if err := c.repository.WriteAnswer(ctx, answeredMessage); err != nil {
		return err
	}

	return nil
}
