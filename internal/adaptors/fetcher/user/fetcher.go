package user

import (
	"context"
)

func NewUserFetcher() *UsrFetcher {
	return &UsrFetcher{
		client: NewClient(),
	}
}

type UsrFetcher struct {
	client *Client
}

func (f *UsrFetcher) Fetch(ctx context.Context, userName string) (u User, err error) {
	u, err = f.client.GetUser(ctx, userName)
	if err != nil {
		return u, err
	}
	return u, nil
}
