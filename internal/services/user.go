package services

import (
	"context"
	"github-user-service/internal/domain/adaptors/fetcher/user"
	"github-user-service/internal/domain/adaptors/streaming"
	"github-user-service/internal/domain/events"
	"github-user-service/internal/domain/models"
	"github-user-service/internal/domain/services"
	"github.com/google/uuid"
)

type UserService struct {
	fetcher  user.Fetcher
	producer streaming.Producer
}

func NewUserService(fetcher user.Fetcher, producer streaming.Producer) services.UserService {
	return &UserService{
		fetcher:  fetcher,
		producer: producer,
	}
}

func (u UserService) GetUser(ctx context.Context, userName string) (id uuid.UUID, err error) {

	usr, err := u.fetcher.Fetch(ctx, userName)
	if err != nil {
		return uuid.UUID{}, err
	}
	followers, err := u.fetcher.FetchFollowers(ctx, userName)

	repos, err := u.fetcher.FetchRepos(ctx, userName)

	followersList := extractFollowers(followers)
	reposList := extractRepos(repos)

	e := createEvent(usr, followersList, reposList)

	err = u.producer.Produce(e)
	if err != nil {
		return uuid.UUID{}, err
	}

	return e.Payload.UniqueID, nil
}

func createEvent(user models.User, followers, repos []string) events.UserChanged {
	meta := models.NewMeta()
	meta.Type = events.UserInfoChangedType

	payload := events.UserPayload{
		UniqueID:  uuid.New(),
		Id:        user.Id,
		Username:  user.Login,
		Followers: followers,
		Repos:     repos,
	}

	return events.UserChanged{
		EventMeta: meta,
		Payload:   payload,
	}
}

func extractFollowers(followers []models.Follower) (followersList []string) {
	followersList = make([]string, 0)
	for _, f := range followers {
		followersList = append(followersList, f.Login)
	}

	return
}

func extractRepos(repos []models.Repository) (reposList []string) {
	reposList = make([]string, 0)
	for _, r := range repos {
		reposList = append(reposList, r.Name)
	}
	return
}
