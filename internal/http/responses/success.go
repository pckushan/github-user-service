package responses

import "github.com/google/uuid"

type Success struct {
	ID uuid.UUID `json:"id"`
}
