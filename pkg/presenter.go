package pkg

type BaseResponse struct {
	StatusCode   int                `json:"status_code"`
	ErrorMessage string             `json:"error_message,omitempty"`
	StackTrace   string             `json:"stack_trace,omitempty"`
	Data         *interface{}       `json:"data"`
	Errors       *[]ErrorValidation `json:"errors"`
	Version      string             `json:"version"`
}

type ErrorValidation struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type PaginationResponse[T any] struct {
	Items *[]T          `json:"items"`
	Meta  *MetaResponse `json:"meta"`
}

type MetaResponse struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}
