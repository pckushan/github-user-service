package fetcher

import (
	"context"
	"github-user-service/internal/domain/models"
)

func NewFetcher() *UsrFetcher {
	return &UsrFetcher{
		client: NewClient(),
	}
}

type UsrFetcher struct {
	client *Client
}

func (f *UsrFetcher) Fetch(ctx context.Context, userName string) (u models.User, err error) {
	usr, err := f.client.GetUser(ctx, userName)
	if err != nil {
		return u, err
	}
	u = mapUser(usr)
	return u, nil
}

func (f *UsrFetcher) FetchFollowers(ctx context.Context, userName string) (followers []models.Follower, err error) {
	followers, err = f.client.GetFollowers(ctx, userName)
	if err != nil {
		return followers, err
	}
	return followers, nil
}

func (f *UsrFetcher) FetchRepos(ctx context.Context, userName string) (repos []models.Repository, err error) {
	repos, err = f.client.GetRepositories(ctx, userName)
	if err != nil {
		return repos, err
	}
	return repos, nil
}

func mapUser(u User) models.User {
	user := models.User{
		Login:           u.Login,
		Id:              u.Id,
		Url:             u.Url,
		FollowersUrl:    u.FollowersUrl,
		FollowingUrl:    u.FollowingUrl,
		ReposUrl:        u.ReposUrl,
		Type:            u.Type,
		SiteAdmin:       u.SiteAdmin,
		Name:            u.Name,
		Company:         u.Company,
		Location:        u.Location,
		Email:           u.Email,
		Bio:             u.Bio,
		TwitterUsername: u.TwitterUsername,
		PublicRepos:     u.PublicRepos,
		Followers:       u.Followers,
		Following:       u.Following,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
	}

	return user
}
