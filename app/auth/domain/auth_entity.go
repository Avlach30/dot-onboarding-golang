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
