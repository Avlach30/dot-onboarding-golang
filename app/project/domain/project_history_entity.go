package domain

import "time"

type ProjectHistoryEntity struct {
	ID            int64
	HistoryType   string
	Title         string
	Description   string
	AttachmentUrl string
	CreatedAt     time.Time
}
