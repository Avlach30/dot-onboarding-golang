package pkg

type BaseResponse struct {
	StatusCode   int                `json:"status_code"`
	ErrorMessage string             `json:"error_message,omitempty"`
	StackTrace   string             `json:"stack_trace"`
	Data         *interface{}       `json:"data"`
	Errors       *[]ErrorValidation `json:"errors"`
	Version      string             `json:"version"`
}

type ErrorValidation struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type PaginationResponse struct {
	Items *[]interface{} `json:"items"`
	Meta  *MetaResponse  `json:"meta,omitempty"`
}

type MetaResponse struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}
