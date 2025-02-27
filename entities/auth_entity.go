package entities

import (
	"time"

	"github.com/google/uuid"
)

type AuthEntity struct {
	ID             uuid.UUID
	Name           string
	Email          string
	ExpirationTime time.Time
}

type AuthPermissionEntity struct {
	ID   uuid.UUID
	Name string
	Key  string
}
