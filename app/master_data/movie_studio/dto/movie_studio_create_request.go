package dto

type MovieStudioCreateRequest struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
	ChairCapacity int `json:"chair_capacity" binding:"required,gte=0"`
	AdditionalCapacities []string `json:"additional_capacities" binding:"required"`
}