package entities

import (
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type TicketEntity struct {
	ID              uuid.UUID           `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	MovieScheduleId uuid.UUID           `gorm:"type:uuid;index:idx_movie_schedule_id" json:"movie_schedule_id"`
	MovieSchedule   MovieScheduleEntity `gorm:"foreignKey:MovieScheduleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"movie_schedule"`
	UserId          uuid.UUID           `gorm:"type:uuid;index:idx_user_id" json:"user_id"`
	User            UserEntity          `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user"`
	SelectedChairs  utils.StringArray   `gorm:"type:jsonb" json:"selected_chairs"`
	Status          string              `gorm:"type:varchar(255);default:confirmed" json:"status"`
	CreatedAt       time.Time           `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt       time.Time           `gorm:"autoUpdateTime:true" json:"updated_at"`
	DeletedAt       gorm.DeletedAt      `gorm:"index" json:"deleted_at,omitempty"`
}

func (TicketEntity) TableName() string {
	return "tickets"
}
