package services

import (
	"context"
	"github.com/google/uuid"
)

type UserService interface {
	GetUser(ctx context.Context, userName string) (id uuid.UUID, err error)
}
