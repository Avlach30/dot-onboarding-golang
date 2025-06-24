package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieEntity struct {
	ID                uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Title             string         `gorm:"type:varchar(255);not null" json:"title"`
	Genre             string         `gorm:"type:varchar(255);not null" json:"genre"`
	PosterUrl         string         `gorm:"type:varchar(255);not null" json:"poster_url"`
	DurationInMinutes int            `gorm:"type:int;not null" json:"duration_in_minutes"`
	Desciption        string         `gorm:"type:varchar(255);not null" json:"desciption"`
	CreatedAt         time.Time      `gorm:"autoCreateTime:true index" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime:true index" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (MovieEntity) TableName() string {
	return "movies"
}
