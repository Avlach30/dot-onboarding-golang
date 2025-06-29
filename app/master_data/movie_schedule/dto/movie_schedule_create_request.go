package dto

import (
	"time"
	"github.com/google/uuid"
)

type MovieScheduleCreateRequest struct {
	MovieID       uuid.UUID `json:"movie_id" binding:"required,uuid"`
	MovieStudioID uuid.UUID `json:"movie_studio_id" binding:"required,uuid"`
	ShowDatetime  time.Time `json:"show_datetime" binding:"required"`
	Price         float64   `json:"price" binding:"required,gte=1"`
}