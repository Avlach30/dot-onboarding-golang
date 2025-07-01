package utils

import (
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gorm.io/gorm"
)

func GetPaginationOffset(page, limit int) int {
	page -= 1
	offset := page * limit
	return offset
}

func Paginate(queryDto *querydto.QueryDto) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := queryDto.Page
		if page <= 0 {
			page = 1
		}

		perPage := queryDto.PerPage
		switch {
		case perPage > 100:
			perPage = 100
		case perPage <= 0:
			perPage = 10
		}

		offset := (page - 1) * perPage
		return db.Offset(offset).Limit(perPage)
	}
}
