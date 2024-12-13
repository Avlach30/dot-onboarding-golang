package pkg

type BaseResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    *interface{}   `json:"data"`
	Errors  *[]interface{} `json:"errors"`
	Meta    *MetaResponse  `json:"meta,omitempty"`
}

type PaginationResponse struct {
	Items *[]interface{} `json:"items"`
	Meta  *MetaResponse  `json:"meta,omitempty"`
}

type MetaResponse struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
