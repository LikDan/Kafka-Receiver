package controller

import (
	"context"
	"receiver/internal/entities"
)

type IController interface {
	Answer(ctx context.Context, message entities.Message) error
}
