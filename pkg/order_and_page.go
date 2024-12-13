package pkg

type Pagination struct {
	Page  int `json:"page" query:"page"`
	Limit int `json:"limit" query:"limit"`
}

type Ordering struct {
	OrderBy   string `json:"order_by" query:"order_by"`
	OrderType string `json:"order_type" query:"order_type" binding:"omitempty,oneof=ASC DESC asc desc"`
}
