package user

import (
	"context"
	"github-user-service/internal/domain/models"
)

type Fetcher interface {
	Fetch(ctx context.Context, userName string) (models.User, error)
	FetchRepos(ctx context.Context, userName string) (repos []models.Repository, err error)
	FetchFollowers(ctx context.Context, userName string) (followers []models.Follower, err error)
}
