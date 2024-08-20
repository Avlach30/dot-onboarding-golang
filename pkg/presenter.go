package pkg

type BaseResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
	Meta    *MetaResponse `json:"meta,omitempty"`
}

type MetaResponse struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
