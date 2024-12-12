package utils

func GetPaginationOffset(page, limit int) int {
	page -= 1
	offset := page * limit
	return offset
}
