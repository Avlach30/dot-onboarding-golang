package domain

import "time"

type Entity struct {
	ID                int64
	UUID              string
	Name              string
	Description       string
	ThumbnailImageURL string
	ServiceType       string
	Status            string
	TargetTime        string
	CreatedAt         time.Time
	Astrodevs         string
}
