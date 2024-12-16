package dto

import "time"

type AuthSignInResponse struct {
	ExpiredAt time.Time `json:"expired_at" binding:"required"`
	Token     string    `json:"token" binding:"required"`
	Type      string    `json:"type" binding:"required"`
}
