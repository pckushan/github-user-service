package models

import "time"

type User struct {
	Login       string      `json:"login"`
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	Company     interface{} `json:"company"`
	Location    string      `json:"location"`
	Email       interface{} `json:"email"`
	PublicRepos int         `json:"public_repos"`
	Followers   int         `json:"followers"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
