package user

import (
	"context"
	"github-user-service/internal/adaptors/fetcher/user"
)

type Fetcher interface {
	Fetch(ctx context.Context, userName string) (user.User, error)
}
