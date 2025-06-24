package entities

import (
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type MovieStudioEntity struct {
	ID                   uuid.UUID         `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name                 string            `gorm:"type:varchar(255);not null" json:"name"`
	ChairCapacity        int               `gorm:"type:int;not null" json:"chair_capacity"`
	AdditionalCapacities utils.StringArray `gorm:"type:jsonb" json:"additional_capacities,omitempty"`
	CreatedAt            time.Time         `gorm:"autoCreateTime:true index" json:"created_at"`
	UpdatedAt            time.Time         `gorm:"autoUpdateTime:true index" json:"updated_at"`
	DeletedAt            gorm.DeletedAt    `gorm:"index" json:"deleted_at,omitempty"`
}

func (MovieStudioEntity) TableName() string {
	return "movie_studios"
}
