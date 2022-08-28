package repository

import "github.com/google/uuid"

type Repo interface {
	Insert(record interface{}) error
	Update(record interface{}) error
	Check(id uuid.UUID) bool
}
