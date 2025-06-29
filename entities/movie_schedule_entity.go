package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MovieScheduleEntity struct {
	ID            uuid.UUID         `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	MovieID       uuid.UUID         `gorm:"type:uuid;index:movie_schedules_movie_id_idx" json:"movie_id"`
	Movie         MovieEntity       `gorm:"foreignKey:MovieID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie"`
	MovieStudioID uuid.UUID         `gorm:"type:uuid;index:movie_schedules_movie_studio_id_idx" json:"movie_studio_id"`
	MovieStudio   MovieStudioEntity `gorm:"foreignKey:MovieStudioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_studio"`
	ShowDatetime  time.Time         `gorm:"type:timestamp" json:"show_datetime"`
	Price         float64           `gorm:"type:float" json:"price"`
	CreatedAt     time.Time         `gorm:"autoCreateTime:true;type:timestamp" json:"created_at"`
	UpdatedAt     time.Time         `gorm:"autoUpdateTime:true;type:timestamp" json:"updated_at"`
	DeletedAt     gorm.DeletedAt    `json:"deleted_at,omitempty"` 
}

func (MovieScheduleEntity) TableName() string {
	return "movie_schedules"
}
