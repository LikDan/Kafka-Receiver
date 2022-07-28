package repository

import (
	"context"
	"receiver/internal/entities"
)

type IRepository interface {
	WriteAnswer(ctx context.Context, message entities.AnsweredMessage) error
}
