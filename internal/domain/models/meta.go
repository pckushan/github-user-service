package models

import (
	"github.com/google/uuid"
	"time"
)

const (
	base10 = 10

	ServiceID = `github-user-service`
)

type Meta struct {
	Type      string    `json:"type"`
	EventID   string    `json:"event_id"`
	CreatedAt int64     `json:"created_at"`
	TraceID   uuid.UUID `json:"trace_id"`
	ServiceID string    `json:"service_id"`
}

func NewMeta() Meta {
	return Meta{
		EventID:   uuid.New().String(),
		CreatedAt: time.Now().UTC().UnixNano() / base10,
		TraceID:   uuid.New(),
		ServiceID: ServiceID,
	}
}
