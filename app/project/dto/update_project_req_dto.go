package dto

type UpdateProjectReq struct {
	ServiceType  int    `json:"service_type"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	DeadlineType int    `json:"deadline_type"`
}
