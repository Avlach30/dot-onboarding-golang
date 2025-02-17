package dto

import "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"

type RoleResponse struct {
	Name string `json:"name" binding:"required"`
	Key  string `json:"key" binding:"required"`
}

func RoleResponseFromEntities(datas []entities.RoleEntity) []RoleResponse {
	roleResponses := make([]RoleResponse, len(datas))
	for i, role := range datas {
		roleResponses[i] = RoleResponse{
			Name: role.Name,
			Key:  role.Key,
		}
	}

	return roleResponses
}
