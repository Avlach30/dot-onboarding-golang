package domain

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

type AuthPermission struct {
	ID   uuid.UUID
	Name string
	Key  string
}
