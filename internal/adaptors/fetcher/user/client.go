package user

import (
	"context"
	"encoding/json"
	"fmt"
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
	req, err := c.composeRequest(ctx, userName)
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

func (c *Client) composeRequest(ctx context.Context, userName string) (r *http.Request, err error) {
	reqURL := fmt.Sprintf("%s%s", githubHost, userName)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
