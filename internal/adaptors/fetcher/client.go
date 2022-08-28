package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github-user-service/internal/domain/models"
	"io"
	"net/http"
)

const githubHost = "https://api.github.com/users/"

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: &http.Client{},
	}
}

func (c *Client) GetUser(ctx context.Context, userName string) (User, error) {
	user := User{}
	req, err := c.composeUserRequest(ctx, userName)
	if err != nil {
		return user, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return user, err
	}

	if res.StatusCode != http.StatusOK {
		return user, err
	}
	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (c *Client) composeUserRequest(ctx context.Context, userName string) (r *http.Request, err error) {
	reqURL := fmt.Sprintf("%s%s", githubHost, userName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) GetFollowers(ctx context.Context, userName string) ([]models.Follower, error) {
	followers := make([]models.Follower, 0)
	req, err := c.composeFollowersRequest(ctx, userName)
	if err != nil {
		return followers, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return followers, err
	}

	if res.StatusCode != http.StatusOK {
		return followers, err
	}
	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		return followers, err
	}

	err = json.Unmarshal(reqBody, &followers)
	if err != nil {
		return followers, err
	}

	return followers, nil
}

func (c *Client) composeFollowersRequest(ctx context.Context, userName string) (r *http.Request, err error) {
	reqURL := fmt.Sprintf("%s%s/followers", githubHost, userName)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) GetRepositories(ctx context.Context, userName string) ([]models.Repository, error) {
	repos := make([]models.Repository, 0)
	req, err := c.composeRepositoriesRequest(ctx, userName)
	if err != nil {
		return repos, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return repos, err
	}

	if res.StatusCode != http.StatusOK {
		return repos, err
	}
	reqBody, err := io.ReadAll(res.Body)
	if err != nil {
		return repos, err
	}

	err = json.Unmarshal(reqBody, &repos)
	if err != nil {
		return repos, err
	}

	return repos, nil
}

func (c *Client) composeRepositoriesRequest(ctx context.Context, userName string) (r *http.Request, err error) {
	reqURL := fmt.Sprintf("%s%s/repos", githubHost, userName)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
