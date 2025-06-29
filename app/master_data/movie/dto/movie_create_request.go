package dto

import "mime/multipart"

type MovieCreateRequest struct {
	Title             string `form:"title" binding:"required,min=3,max=255"`
	Genre             string `form:"genre" binding:"required,min=3,max=255"`
	DurationInMinutes int    `form:"duration_in_minutes" binding:"required,gte=0"`
	Description        string `form:"description" binding:"required,min=3,max=255"`
	Poster            multipart.FileHeader `form:"poster" binding:"required"`
}