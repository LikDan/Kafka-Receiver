package controller

import (
	"context"
	"receiver/internal/entities"
)

type IController interface {
	Answer(ctx context.Context, message entities.Message, answer string) error
	Read(ctx context.Context) (entities.Message, error)
}
