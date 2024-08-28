package dto

import userdto "github.com/codespace-id/codespace-x/app/user/dto"

type ProjectDetailResponse struct {
	UUID              string                       `json:"uuid"`
	Name              string                       `json:"name"`
	Description       string                       `json:"description"`
	ServiceType       string                       `json:"service_type"`
	ThumbnailImageURL string                       `json:"thumbnail_image_url"`
	Status            string                       `json:"status"`
	CreatedAt         string                       `json:"created_at"`
	Astrodevs         []userdto.GetProfileResponse `json:"astrodevs"`
}
