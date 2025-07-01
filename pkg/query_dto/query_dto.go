package querydto

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type QueryDto struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Search  string `json:"search"`
	SortBy  string `json:"sort_by"`
	Order   string `json:"order"`
}

func AssignFromHttpContext(httpContext *gin.Context) *QueryDto {
	queryDto := &QueryDto{}
	queryDto.Page, _ = strconv.Atoi(httpContext.DefaultQuery("page", "1"))
	queryDto.PerPage, _ = strconv.Atoi(httpContext.DefaultQuery("per_page", "10"))
	queryDto.Search = httpContext.Query("search")
	queryDto.SortBy = httpContext.Query("sort_by")
	queryDto.Order = httpContext.DefaultQuery("order", "desc")

	return queryDto
}