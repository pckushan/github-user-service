package events

import (
	"github-user-service/internal/domain/models"
	"github.com/google/uuid"
)

const UserInfoChangedType = "UserInfoChanged"

type UserChanged struct {
	EventMeta models.Meta `json:"meta"`
	Payload   UserPayload `json:"payload"`
}

type UserPayload struct {
	UniqueID   uuid.UUID `json:"unique_id"`
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Followers  []string  `json:"followers"`
	Repos      []string  `json:"repos"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	TimeZoneId string    `json:"time_zone_id"`
}
