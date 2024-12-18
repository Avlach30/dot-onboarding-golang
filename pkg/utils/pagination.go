package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPaginationOffset(page, limit int) int {
	page -= 1
	offset := page * limit
	return offset
}

func Paginate(ctx *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageStr := ctx.DefaultQuery("page", "1")
		page, _ := strconv.Atoi(pageStr)
		if page <= 0 {
			page = 1
		}

		perPageStr := ctx.DefaultQuery("per_page", "10")
		perPage, _ := strconv.Atoi(perPageStr)
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
